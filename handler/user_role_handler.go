package handler

// import (
// 	"net/http"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/user_app/model"
// 	"github.com/user_app/repository"
// )

// type UserRoleHandler struct {
// 	roleRepo *repository.RoleRepository
// }

// func NewUserRoleHandler(roleRepo *repository.RoleRepository) *UserRoleHandler {
// 	return &UserRoleHandler{
// 		roleRepo: roleRepo,
// 	}
// }

// func (u *UserRoleHandler) GetUserRoles(c *fiber.Ctx) error {

// 	var userRoles []model.Role

// 	err := u.roleRepo.GetRoles(&userRoles)

// 	if err != nil {
// 		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{"status": "error", "message": "Error on get user roles", "data": err})
// 	}

// 	return c.Status(http.StatusOK).JSON(&fiber.Map{"status": "success", "message": "User roles found", "data": userRoles})

// }

// func (u *UserRoleHandler) GetUserRole(c *fiber.Ctx) error {

// 	id, err := c.ParamsInt("id")

// 	if err != nil {
// 		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{"status": "error", "message": "Id is invalid", "data": err})
// 	}

// 	userRole, erre := u.roleRepo.GetRole(id)

// 	if erre != nil {
// 		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{"status": "error", "message": "Error on get user role", "data": err})
// 	}

// 	return c.Status(http.StatusOK).JSON(&fiber.Map{"status": "success", "message": "User role found", "data": userRole})

// }

// func (u *UserRoleHandler) CreateUserRole(c *fiber.Ctx) error {

// 	var userRole model.Role

// 	err := c.BodyParser(&userRole)

// 	if err != nil {
// 		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{"status": "error", "message": "Error on request", "data": err})
// 	}

// 	err = u.roleRepo.CreateRole(&userRole)

// 	if err != nil {
// 		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{"status": "error", "message": "Error on create user role", "data": err})
// 	}

// 	return c.Status(http.StatusOK).JSON(&fiber.Map{"status": "success", "message": "User role created", "data": userRole})

// }

// func (u *UserRoleHandler) UpdateUserRole(c *fiber.Ctx) error {

// 	id, err := c.ParamsInt("id")

// 	if err != nil {
// 		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{"status": "error", "message": "Id is invalid", "data": err})
// 	}

// 	var userRole model.Role

// 	err = c.BodyParser(&userRole)

// 	if err != nil {
// 		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{"status": "error", "message": "Error on request", "data": err})
// 	}

// 	err = u.roleRepo.UpdateRole(&userRole, uint(id))

// 	if err != nil {
// 		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{"status": "error", "message": "Error on update user role", "data": err})
// 	}

// 	return c.Status(http.StatusOK).JSON(&fiber.Map{"status": "success", "message": "User role updated", "data": userRole})

// }
