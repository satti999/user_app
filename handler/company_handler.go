package handler

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/user_app/model"
	"github.com/user_app/repository"
	"github.com/user_app/utils"
)

type CompanyHandlerInterface interface {
	CreateCompany(c *fiber.Ctx) error
	GetCompanyByID(c *fiber.Ctx) error
	GetAllCompanies(c *fiber.Ctx) error
	UpdateCompany(c *fiber.Ctx) error
	DeleteCompany(c *fiber.Ctx) error
	GetCompanyByName(c *fiber.Ctx) error
}

type CompanyHandler struct {
	crepo *repository.CompanyRepository
}

func NewCompanyHandler(companyrepo *repository.CompanyRepository) CompanyHandlerInterface {
	return &CompanyHandler{
		crepo: companyrepo,
	}
}

func (ch *CompanyHandler) CreateCompany(c *fiber.Ctx) error {

	id := c.Locals("userID")
	userID, ok := id.(uint)
	if !ok {
		// Handle the error case when the type assertion fails
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get user ID",
		})
	}

	var company model.Company

	err := c.BodyParser(&company)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"status": "error", "message": "Error on request", "data": err})

	}
	result, _ := ch.crepo.CompanyAlreadyExist(company.Name)
	fmt.Println("Reqest company name", company.Name)
	fmt.Println("Response company name", result.Name)

	if result.Name == company.Name {

		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"status": "error", "message": "Can not create company on same name "})
	}
	company.UserID = userID
	result, err = ch.crepo.CreateCompany(&company)
	fmt.Println("company id", result.ID)

	if err != nil {

		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"status": "error", "message": "Error on creating company", "data": err})
	}

	return c.Status(http.StatusCreated).JSON(
		&fiber.Map{"status": "success", "message": "Company created successfully", "company": result, "success": true})

}

func (ch *CompanyHandler) GetCompanyByID(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")

	if err != nil {

		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"status": "error", "message": "Id is invalid", "data": err})

	}

	company, err := ch.crepo.GetCompanyByID(uint(id))

	if err != nil {

		return c.Status(http.StatusNotFound).JSON(
			&fiber.Map{"status": "error", "message": "Company not found", "data": err})

	}

	return c.Status(http.StatusOK).JSON(
		&fiber.Map{"status": "success", "message": "Company found", "company": company, "success": true})

}

func (ch *CompanyHandler) GetAllCompanies(c *fiber.Ctx) error {

	companies, err := ch.crepo.GetAllCompanies()

	if err != nil {

		return c.Status(http.StatusNotFound).JSON(
			&fiber.Map{"status": "error", "message": "Companies not found", "data": err})

	}

	return c.Status(http.StatusOK).JSON(
		&fiber.Map{"status": "success", "message": "Companies found", "companies": companies, "success": true})

}

func (ch *CompanyHandler) UpdateCompany(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	imagUrl := utils.GetProfileUrl()
	if err != nil {

		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"status": "error", "message": "Id is invalid", "data": err})

	}
	ide := c.Locals("userID")
	userID, ok := ide.(uint)
	if !ok {
		// Handle the error case when the type assertion fails
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get user ID",
		})
	}

	var company model.Company

	if err := c.BodyParser(&company); err != nil {

		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"status": "error", "message": "Error on request", "data": err})

	}
	fmt.Println("image url of logo ,in company", imagUrl)
	company.Logo = imagUrl
	company.UserID = userID

	err = ch.crepo.UpdateCompany(company, uint(id))

	if err != nil {

		return c.Status(http.StatusNotFound).JSON(
			&fiber.Map{"status": "error", "message": "Company not found", "data": err})

	}

	return c.Status(http.StatusOK).JSON(
		&fiber.Map{"status": "success", "message": "Company info updated successfully", "success": true})

}

func (ch *CompanyHandler) DeleteCompany(c *fiber.Ctx) error {

	var company model.Company
	err := c.BodyParser(&company)
	if err != nil {

		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"status": "error", "message": "Id is invalid", "data": err})

	}

	err = ch.crepo.DeleteCompany(company)

	if err != nil {

		return c.Status(http.StatusNotFound).JSON(
			&fiber.Map{"status": "error", "message": "Company not found", "data": err})

	}

	return c.Status(http.StatusOK).JSON(
		&fiber.Map{"status": "success", "message": "Company deleted successfully", "data": nil})

}

func (ch *CompanyHandler) GetCompanyByName(c *fiber.Ctx) error {

	name := c.Params("name")

	company, err := ch.crepo.GetCompanyByName(name)

	if err != nil {

		return c.Status(http.StatusNotFound).JSON(
			&fiber.Map{"status": "error", "message": "Company not found", "data": err})

	}

	return c.Status(http.StatusOK).JSON(
		&fiber.Map{"status": "success", "message": "Company found", "data": company})

}
