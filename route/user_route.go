package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/user_app/handler"
	"github.com/user_app/middleware"
	"github.com/user_app/utils"
)

type UserRoute struct {
	userHandler handler.UserHandlerInterface
}

func NewUserRoute(userHandler handler.UserHandlerInterface) *UserRoute {
	return &UserRoute{
		userHandler: userHandler,
	}
}

func (ur *UserRoute) UserRoute(router fiber.Router, app *fiber.App) {
	router.Post("/create", utils.UploadProfileFiles, ur.userHandler.CreateUser)
	router.Post("/login", ur.userHandler.LoginHandler)
	router.Get("/logout", ur.userHandler.Logout)
	router.Get("/google_login", ur.userHandler.GoogleSignin)
	router.Get("/oauth/google/callback", ur.userHandler.GoogleCallback)
	app.Use(middleware.AdminMiddleware)
	router.Get("/get/:id", ur.userHandler.GetUserByID)
	// router.Get("/get/:email/GetUserByEmail", ur.userHandler.GetUserByEmail)
	router.Get("/get", ur.userHandler.GetAllUsers)
	router.Put("/update/:id", utils.UploadProfileFiles, ur.userHandler.UpdateUser)
	router.Delete("/delete/:id", ur.userHandler.DeleteUser)
	router.Put("role", ur.userHandler.UpdateUserRole)
	router.Get("/role/:role", ur.userHandler.GetUserByRole)
}
