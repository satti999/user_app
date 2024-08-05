package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/user_app/handler"
	"github.com/user_app/middleware"
)

type UserRoute struct {
	userHandler *handler.UserHandler
}

func NewUserRoute(userHandler *handler.UserHandler) *UserRoute {
	return &UserRoute{
		userHandler: userHandler,
	}
}

func (ur *UserRoute) UserRoute(app *fiber.App) {
	app.Post("/user", ur.userHandler.CreateUser)
	app.Post("/login", ur.userHandler.LoginHandler)
	app.Use(middleware.AuthMiddleware)
	app.Get("/user/:id", ur.userHandler.GetUserByID)
	app.Get("/user/email/:email", ur.userHandler.GetUserByEmail)
	app.Get("/users", ur.userHandler.GetAllUsers)
	app.Put("/user/:id", ur.userHandler.UpdateUser)
	app.Delete("/user/:id", ur.userHandler.DeleteUser)
	app.Get("/user/role/:role", ur.userHandler.GetUserByRole)
}
