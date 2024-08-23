package utils

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gofiber/fiber/v2"
)

var image_url string
var resume_url string

func generateRandomString() string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 5)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
func Credentials() *cloudinary.Cloudinary {
	cld, err := cloudinary.NewFromParams("dajvb9xb4", "944924327512427", "gbWsNKIG28gr1budOi7-CKgUzJE")
	if err != nil {
		log.Fatalf("Failed to initialize Cloudinary: %s", err)
	}
	return cld
}

func UploadFileToCloudinary(cld *cloudinary.Cloudinary, file multipart.File, folder string) (string, error) {
	ctx := context.Background()
	uploadParams := uploader.UploadParams{
		Folder: folder, // Use the provided folder name
	}
	uploadResult, err := cld.Upload.Upload(ctx, file, uploadParams)
	if err != nil {
		return "", fmt.Errorf("upload error: %v", err)
	}
	return uploadResult.SecureURL, nil
}
func HandleFileUpload(c *fiber.Ctx, fieldName string, folder string) (string, error) {
	file, err := c.FormFile(fieldName)
	if err != nil {
		return "", nil // No file uploaded
	}

	if file.Filename != "" {
		dir, err := os.Getwd()
		if err != nil {
			return "", err
		}

		cld := Credentials()
		tempDir := filepath.Join(dir, "uploads", folder)
		err = os.MkdirAll(tempDir, os.ModePerm)
		if err != nil {
			return "", err
		}

		tempFile, err := file.Open()
		if err != nil {
			return "", fmt.Errorf("failed to open file: %s", err)
		}
		defer tempFile.Close()

		ext := filepath.Ext(file.Filename)
		filename := fmt.Sprintf("%s-%s%s", folder, generateRandomString(), ext)
		tempPath := filepath.Join(tempDir, filename)

		err = c.SaveFile(file, tempPath)
		if err != nil {
			return "", err
		}

		// Upload to Cloudinary
		url, err := UploadFileToCloudinary(cld, tempFile, folder)
		if err != nil {
			return "", err
		}

		// Cleanup
		os.Remove(tempPath)

		return url, nil
	}

	return "", nil
}
func UploadProfileFiles(c *fiber.Ctx) error {

	fmt.Println("************* UploadProfileFiles  ************")
	pfile, err := c.FormFile("image")
	fmt.Println("User in image function")
	if pfile == nil && err != nil {
		rfile, err := c.FormFile("resume")
		if rfile == nil && err != nil {
			c.Next()
		} else {
			imageURL, err := HandleFileUpload(c, "image", "images")
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("failed to upload image: %s", err))

			}
			image_url = imageURL
			return c.Next()
		}

	}
	rfile, err := c.FormFile("resume")
	fmt.Println("User in resume function")
	if rfile == nil && err != nil {
		pfile, err := c.FormFile("image")
		if pfile == nil && err != nil {
			c.Next()
		} else {
			imageURL, err := HandleFileUpload(c, "image", "images")
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("failed to upload image: %s", err))

			}
			image_url = imageURL
			return c.Next()
		}
	}

	if pfile != nil && rfile != nil {
		fmt.Println("User in resume and image function function")
		imageURL, err := HandleFileUpload(c, "image", "images")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("failed to upload image: %s", err))
		}
		image_url = imageURL
		resumeURL, err := HandleFileUpload(c, "resume", "resumes")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("failed to upload resume: %s", err))
		}
		resume_url = resumeURL

		// Set the URLs in the context or return them as part of the response
		fmt.Println("Image URL:", image_url)
		fmt.Println("Resume URL:", resume_url)
		return c.Next()
	}
	return c.Next()

}
func UpdateUserProfile(c *fiber.Ctx) error {

	err := UploadProfileFiles(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to upload files")
	}
	return c.Next()
}

func GetProfileUrl() string {
	return image_url
}

func GetResumeUrl() string {
	return resume_url
}
