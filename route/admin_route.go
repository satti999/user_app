package route

// import (
// 	"github.com/gofiber/fiber/v2"
// 	"github.com/user_app/handler"
// 	"github.com/user_app/middleware"
// )

// type AdminRoute struct {
// 	adHandler *handler.UserRoleHandler
// }

// func NewAdminRoute(admHandler *handler.UserRoleHandler) *AdminRoute {
// 	return &AdminRoute{
// 		adHandler: admHandler,
// 	}
// }

// func (a *AdminRoute) AdminRoute(app *fiber.App) {

// 	app.Post("/admin", a.adHandler.CreateUserRole)

// 	app.Use(middleware.AuthMiddleware)

// 	app.Get("/admin", a.adHandler.GetUserRoles)

// 	// app.Delete("/admin/:id", a.adHandler.)

// 	app.Get("/admin/:id", a.adHandler.GetUserRole)

// 	app.Put("/admin/:id", a.adHandler.UpdateUserRole)

// }
