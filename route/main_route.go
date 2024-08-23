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
	companyRepo := repository.NewCompanyRepository(repo)
	applicationRepo := repository.NewApplicationRepository(repo)
	jobRepo := repository.NewJobRepository(repo)

	//Handlers
	userHandler := handler.NewUserHandler(userRepo)
	companyHandler := handler.NewCompanyHandler(companyRepo)
	applicationHandler := handler.NewApplicationHandler(applicationRepo)
	jobHandler := handler.NewJobHandler(jobRepo)
	// adminHandler := handler.NewUserRoleHandler(adRepo)

	//Routes
	userRoute := NewUserRoute(userHandler)
	companyRoute := NewCompanyRoute(companyHandler)
	applicationRoute := NewApplicationRoute(applicationHandler)
	jobRoute := NewJobRoute(jobHandler)

	//Main Route function
	userGroup := app.Group("/api/v1/user")
	userRoute.UserRoute(userGroup, app)

	companyGroup := app.Group("/api/v1/company")
	companyRoute.CompanyRoute(companyGroup, app)

	jobGroup := app.Group("/api/v1/job")
	jobRoute.JobRoute(jobGroup, app)

	applicationGroup := app.Group("/api/v1/application")
	applicationRoute.Job_Application_Routerouter(applicationGroup, app)

}
