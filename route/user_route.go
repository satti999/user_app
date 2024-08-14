package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/user_app/handler"
	"github.com/user_app/middleware"
)

type UserRoute struct {
	userHandler handler.UserHandlerInterface
}

func NewUserRoute(userHandler handler.UserHandlerInterface) *UserRoute {
	return &UserRoute{
		userHandler: userHandler,
	}
}

func (ur *UserRoute) UserRoute(app *fiber.App) {
	app.Post("/user", ur.userHandler.CreateUser)
	app.Post("/login", ur.userHandler.LoginHandler)
	app.Get("/google_login", ur.userHandler.GoogleSignin)
	app.Get("/oauth/google/callback", ur.userHandler.GoogleCallback)
	app.Use(middleware.AuthMiddleware, middleware.AdminMiddleware)
	app.Get("/user/:id", ur.userHandler.GetUserByID)
	app.Get("/user/:email/GetUserByEmail", ur.userHandler.GetUserByEmail)
	app.Get("/users", ur.userHandler.GetAllUsers)
	app.Put("/user/:id", ur.userHandler.UpdateUser)
	app.Delete("/user/:id", ur.userHandler.DeleteUser)
	app.Get("/user/role/:role", ur.userHandler.GetUserByRole)
}
