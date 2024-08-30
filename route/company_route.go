package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/user_app/handler"
	"github.com/user_app/middleware"
	"github.com/user_app/utils"
)

type CompanyRoute struct {
	companyHandler handler.CompanyHandlerInterface
}

func NewCompanyRoute(companyHandler handler.CompanyHandlerInterface) *CompanyRoute {
	return &CompanyRoute{
		companyHandler: companyHandler,
	}
}

func (cr *CompanyRoute) CompanyRoute(router fiber.Router, app *fiber.App) {

	app.Use(middleware.AdminMiddleware)
	router.Post("/register", cr.companyHandler.CreateCompany)
	router.Get("/get/:id", cr.companyHandler.GetCompanyByID)
	router.Get("/get", cr.companyHandler.GetAllCompanies)
	router.Put("/update/:id", utils.UploadProfileFiles, cr.companyHandler.UpdateCompany)
	router.Delete("/delete/:id", cr.companyHandler.DeleteCompany)
	router.Get("/get/:name", cr.companyHandler.GetCompanyByName)

}
