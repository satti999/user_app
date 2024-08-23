package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/user_app/handler"
	"github.com/user_app/middleware"
	"github.com/user_app/utils"
)

type ApplicationRoute struct {
	applicationHandler *handler.ApplicationHandler
}

func NewApplicationRoute(applicationHandler *handler.ApplicationHandler) *ApplicationRoute {
	return &ApplicationRoute{
		applicationHandler: applicationHandler,
	}
}

func (ar *ApplicationRoute) Job_Application_Routerouter(router fiber.Router, app *fiber.App) {

	app.Use(middleware.AuthMiddleware, middleware.AdminMiddleware)

	router.Post("/create", ar.applicationHandler.ApplyForJob)

	router.Get("/get", ar.applicationHandler.GetAppliedJobs)
	router.Put("/update/:id", utils.UpdateUserProfile, ar.applicationHandler.UpdateStatus)
	// router.Delete("/delete/:id", ar.applicationHandler.DeleteApplication)

}
