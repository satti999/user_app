package handler

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/user_app/model"
	"github.com/user_app/repository"
)

type ApplicationHandler struct {
	arepo *repository.ApplicationRepository
}

func NewApplicationHandler(aplirepo *repository.ApplicationRepository) *ApplicationHandler {
	return &ApplicationHandler{
		arepo: aplirepo,
	}
}

func (h *ApplicationHandler) ApplyForJob(c *fiber.Ctx) error {
	fmt.Println("user applied for job")
	application := model.Application{}
	id := c.Locals("userID")
	userID, ok := id.(uint)
	if !ok {
		// Handle the error case when the type assertion fails
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get user ID",
		})
	}
	jobid, err := c.ParamsInt("id")

	if err != nil {

		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"status": "error", "message": "Id is invalid", "data": err})

	}

	// if err := c.BodyParser(&application); err != nil {
	// 	return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
	// 		"status":  "fail",
	// 		"message": err.Error(),
	// 	})
	// }

	application.UserID = userID
	application.JobID = uint(jobid)

	err = h.arepo.ApplyJob(&application)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"status":      "success",
		"message":     "Job applied successfully",
		"application": application,
		"success":     true,
	})

}

func (h *ApplicationHandler) GetAppliedJobs(c *fiber.Ctx) error {

	id := c.Locals("userID")
	userID, ok := id.(uint)
	if !ok {
		// Handle the error case when the type assertion fails
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get user ID",
		})
	}
	application, err := h.arepo.GetAppliedJobs(userID)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{"status": "error", "message": "Error on request", "data": err})
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{"status": "success", "message": "Application found", "application": application, "success": true})

}

func (h *ApplicationHandler) UpdateStatus(c *fiber.Ctx) error {
	applicationid, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{"status": "error", "message": "Id is invalid", "data": nil})
	}

	var application model.Application

	if err := c.BodyParser(&application); err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{"status": "error", "message": "Error on request", "data": err})
	}

	err = h.arepo.UpdateStatus(string(application.Status), uint(applicationid))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{"status": "error", "message": "Error on request", "data": err})
	}

	return c.Status(http.StatusOK).JSON(&fiber.Map{"status": "success", "message": "Status updated successfully", "application": application, "success": true})

}

func (h *ApplicationHandler) GetApplicationByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{"status": "error", "message": "Id is invalid", "data": nil})
	}

	application, err := h.arepo.GetApplication(uint(id))

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{"status": "error", "message": "Error on request", "data": err})
	}

	return c.Status(http.StatusOK).JSON(&fiber.Map{"status": "success", "message": "Application found", "application": application, "success": true})

}
