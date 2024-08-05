package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	database "github.com/user_app/database"
	"github.com/user_app/route"
)

func main() {
	app := fiber.New()
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	configration := &database.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}
	db := database.ConnectDB(configration)

	database.Migrate(db)

	route.MainRoute(app, db)

	err = app.Listen(":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

}
