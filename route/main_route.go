package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/user_app/handler"
	"github.com/user_app/repository"
	"gorm.io/gorm"
)

func MainRoute(app *fiber.App, db *gorm.DB) {
	repo := repository.NewRepository(db)
	userRepo := repository.NewUserRepository(repo)
	userHandler := handler.NewUserHandler(userRepo)

	userRoute := NewUserRoute(userHandler)

	userRoute.UserRoute(app)

}
