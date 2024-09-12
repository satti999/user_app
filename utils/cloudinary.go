package utils

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gofiber/fiber/v2"
)

var image_url string
var resume_url string
var file_name string

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

func UploadFileToCloudinary(cld *cloudinary.Cloudinary, file multipart.File, folder string, num int) (string, error) {
	ctx := context.Background()
	if num == 0 {
		uploadParams := uploader.UploadParams{
			Folder: folder, // Use the provided folder name
		}
		uploadResult, err := cld.Upload.Upload(ctx, file, uploadParams)
		if err != nil {
			return "", fmt.Errorf("upload error: %v", err)
		}
		fmt.Println("upload image response", uploadResult.OriginalFilename)
		return uploadResult.SecureURL, nil
	}
	uploadParams := uploader.UploadParams{
		Folder:       folder, // Use the provided folder name
		ResourceType: "raw",
		Format:       "pdf",
	}
	uploadResult, err := cld.Upload.Upload(ctx, file, uploadParams)
	if err != nil {
		return "", fmt.Errorf("upload error: %v", err)
	}
	fmt.Println("upload resume file response", uploadResult.OriginalFilename)
	fmt.Println("resume fie url ", uploadResult.SecureURL)
	return uploadResult.SecureURL, nil

}
func HandleFileUpload(c *fiber.Ctx, fieldName string, folder string, num int) (string, error) {
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
		if num == 1 {
			file_name = file.Filename
		}

		ext := filepath.Ext(file.Filename)
		filename := fmt.Sprintf("%s-%s%s", folder, generateRandomString(), ext)
		tempPath := filepath.Join(tempDir, filename)

		err = c.SaveFile(file, tempPath)
		if err != nil {
			return "", err
		}

		// Upload to Cloudinary
		url, err := UploadFileToCloudinary(cld, tempFile, folder, num)
		if err != nil {
			return "", err
		}

		// Cleanup
		os.Remove(tempPath)

		return url, nil
	}

	return "", nil
}
func UploadImage(c *fiber.Ctx) error {
	fmt.Println("image called in cloudinary")
	pfile, err := c.FormFile("image")

	if pfile == nil && err != nil {
		fmt.Println("image url------- in upload image")
		return c.Next()
	}

	imageURL, err := HandleFileUpload(c, "image", "images", 0)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("failed to upload image: %s", err))

	}

	image_url = imageURL

	return c.Next()

}
func UploadResume(c *fiber.Ctx) error {

	rfile, err := c.FormFile("resume")

	if rfile == nil && err != nil {
		return c.Next()
	}

	resumeURL, err := HandleFileUpload(c, "resume", "resumes", 1)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("failed to upload image: %s", err))

	}

	resume_url = resumeURL

	return c.Next()
}

func GetProfileUrl() string {
	return image_url
}

func GetResumeUrl() string {
	return resume_url
}
func GetFileName() string {
	baseName := strings.Split(file_name, ".")[0]
	return baseName
}
