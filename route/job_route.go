package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/user_app/handler"
	"github.com/user_app/middleware"
)

type JobRoute struct {
	jobHandler *handler.JobHandler
}

func NewJobRoute(jobHandler *handler.JobHandler) *JobRoute {
	return &JobRoute{
		jobHandler: jobHandler,
	}
}

func (jr *JobRoute) JobRoute(router fiber.Router, app *fiber.App) {

	app.Use(middleware.AdminMiddleware)
	router.Post("/post", jr.jobHandler.PostJob)
	router.Get("/get/:id", jr.jobHandler.GetJobByID)
	router.Get("/get", jr.jobHandler.GetAllJobs)
	router.Get("/getadminjobs", jr.jobHandler.GetAdminJobs)
	router.Put("/update/:id", jr.jobHandler.UpdateJob)
	router.Delete("/delete/:id", jr.jobHandler.DeleteJob)

}
