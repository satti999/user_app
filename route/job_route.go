package route
import (
	"github.com/gofiber/fiber/v2"
	"github.com/user_app/handler"
	"github.com/user_app/middleware"
	"github.com/user_app/utils"
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

	app.Use(middleware.AuthMiddleware, middleware.AdminMiddleware)
	router.Post("/create", jr.jobHandler.PostJob)
	router.Get("/get/:id", jr.jobHandler.GetJobByID)
	router.Get("/get", jr.jobHandler.GetAllJobs)
	router.Put("/update/:id",utils.UploadProfileFiles, jr.jobHandler.UpdateJob)
	router.Delete("/delete/:id", jr.jobHandler.DeleteJob)
	

}

