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
	jobRepo:= repository.NewJobRepository(repo)
	// adRepo := repository.NewRoleRepository(repo)
	//Handlers
	userHandler := handler.NewUserHandler(userRepo)
	companyHandler := handler.NewCompanyHandler(companyRepo)
	jobHandler := handler.NewJobHandler(jobRepo)
	// adminHandler := handler.NewUserRoleHandler(adRepo)

	//Routes
	userRoute := NewUserRoute(userHandler)
	companyRoute := NewCompanyRoute(companyHandler)
	jobRoute := NewJobRoute(jobHandler)

	//Main Route function
	userGroup := app.Group("/api/v1/user")
	userRoute.UserRoute(userGroup,app)

	companyGroup:=app.Group("/api/v1/company")
	companyRoute.CompanyRoute(companyGroup,app)

	jobGroup:= app.Group("/api/v1/job")
	jobRoute.JobRoute(jobGroup,app)

	

}
