package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/user_app/handler"
	"github.com/user_app/repository"
	"gorm.io/gorm"
)

func MainRoute(app *fiber.App, db *gorm.DB) {
	repo := repository.NewRepository(db)
	//Repos
	userRepo := repository.NewUserRepository(repo)
	// adRepo := repository.NewRoleRepository(repo)
	//Handlers
	userHandler := handler.NewUserHandler(userRepo)
	// adminHandler := handler.NewUserRoleHandler(adRepo)

	//Routes
	userRoute := NewUserRoute(userHandler)
	// adminRoute := NewAdminRoute(adminHandler)

	//Main Route function
	userRoute.UserRoute(app)
	// adminRoute.AdminRoute(app)

}
