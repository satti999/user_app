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

	app.Use(middleware.AdminMiddleware)

	router.Get("/create/:id", ar.applicationHandler.ApplyForJob)

	router.Get("/get", ar.applicationHandler.GetAppliedJobs)
	router.Put("/update/:id", utils.UploadResume, ar.applicationHandler.UpdateStatus)
	router.Get("/:id/applicants", ar.applicationHandler.GetApplicationByID)
	// router.Delete("/delete/:id", ar.applicationHandler.DeleteApplication)

}
