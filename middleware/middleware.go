package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/user_app/model"
)

var AdminRole = "admin"

func CreateToken(user model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": user.Name,
			"id":       user.ID,
			"email":    user.Email,
			"role":     user.Role,
			"exp":      time.Now().Add(time.Minute * 30).Unix(),
		})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func AuthMiddleware(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	if cookie == "" {
		return c.Status(http.StatusUnauthorized).JSON(&fiber.Map{"error": "missing token"})
	}
	// if authHeader == "" {
	// 	return c.Status(http.StatusUnauthorized).JSON(&fiber.Map{"error": "missing token"})
	// }
	//tokenString := cookie[len(""):]
	//fmt.Println("token String", tokenString)
	_, err := jwt.ParseWithClaims(cookie, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return err
	}

	return c.Next()
}

func AdminMiddleware(c *fiber.Ctx) error {

	// var userRole string
	cookie := c.Cookies("jwt")

	if cookie == "" {
		return c.Status(http.StatusUnauthorized).JSON(&fiber.Map{"error": "missing token "})

	}
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(cookie, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return err
	}

	userRole, ok := claims["role"].(string)
	if !ok {
		panic("Couldn't parse email as string")
	}
	fmt.Println(userRole)

	if userRole != "admin" {
		return c.Status(http.StatusUnauthorized).JSON(&fiber.Map{"error": "Access denied as only admin allowed"})
	}
	return c.Next()
}
