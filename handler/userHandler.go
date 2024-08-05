package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/user_app/middleware"
	"github.com/user_app/model"
	"github.com/user_app/repository"
	"github.com/user_app/utils"
)

type UserHandler struct {
	urepo *repository.UserRepository
}

func NewUserHandler(useRrepo *repository.UserRepository) *UserHandler {
	return &UserHandler{
		urepo: useRrepo,
	}
}

func (uh *UserHandler) CreateUser(c *fiber.Ctx) error {

	User := model.User{}

	err := c.BodyParser(&User)
	fmt.Println("user", User.Name)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{"status": "error", "message": "Error on request", "data": err})

	}

	exists := uh.urepo.UserExists(User.Email)

	if exists {
		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "User already exists"})
	}
	pass := User.Password
	hashedPass := utils.HashAndSalt(pass)
	User.Password = hashedPass

	err = uh.urepo.CreateUser(User)

	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Error on creating user"})
		return err

	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "User has been created successfully"})
	return nil

}
func (uh *UserHandler) LoginHandler(c *fiber.Ctx) error {

	var u model.User

	err := c.BodyParser(&u)

	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Error on login request"})
		return err
	}

	user, err := uh.urepo.GetUserByEmail(u.Email)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{"status": "error", "message": "user not found", "data": err})
	}
	err = utils.CheckHash(user.Password, u.Password)
	if err != nil {
		fmt.Println(err)
		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"status": "error", "message": "password not matches", "data": err})
	}

	token, err := middleware.CreateToken(user)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"status": "error", "message": "Error on login request", "data": err})
	}
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{"status": "success",
		"message": "User logged in",
		"data":    token,
		"userId":  user.ID,
		"role":    user.Role,
		"email":   user.Email})

}

func (uh *UserHandler) GetUserByID(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{"status": "error", "message": "Id is invalid", "data": nil})
	}

	user, err := uh.urepo.GetUserByID(uint(id))

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{"status": "error", "message": "Error on getting user", "data": err})
	}

	return c.Status(http.StatusOK).JSON(&fiber.Map{"status": "success", "message": "User found", "data": user})

}

func (uh *UserHandler) GetUserByEmail(c *fiber.Ctx) error {

	email := c.Params("email")

	user, err := uh.urepo.GetUserByEmail(email)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"status": "error", "message": "Error on getting user", "data": err})
	}

	return c.Status(http.StatusOK).JSON(&fiber.Map{"status": "success", "message": "User found", "data": user})

}

func (uh *UserHandler) GetAllUsers(c *fiber.Ctx) error {

	users, err := uh.urepo.GetAllUsers()

	if err != nil {

		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"status": "error", "message": "Error on getting users", "data": err})
	}

	return c.Status(http.StatusOK).JSON(&fiber.Map{"status": "success", "message": "Users found", "data": users})

}

func (uh *UserHandler) UpdateUser(c *fiber.Ctx) error {
	ide := c.Params("id")

	userID, err := strconv.Atoi(ide)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{"status": "error", "message": "Id is invalid", "data": err})
	}
	var user model.User

	user_ID := uint(userID)

	if err := c.BodyParser(&user); err != nil {

		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"status": "error", "message": "Error on request", "data": err})

	}

	err = uh.urepo.UpdateUser(user, user_ID)

	if err != nil {

		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"status": "error", "message": "Error on updating user", "data": err})

	}
	user.ID = user_ID

	return c.Status(http.StatusOK).JSON(&fiber.Map{"status": "success", "message": "User updated", "data": user})

}

func (uh *UserHandler) DeleteUser(c *fiber.Ctx) error {

	user := new(model.User)

	if err := c.BodyParser(&user); err != nil {

		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"status": "error", "message": "Error on request", "data": err})

	}

	err := uh.urepo.DeleteUser(*user)

	if err != nil {

		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"status": "error", "message": "Error on deleting user", "data": err})

	}

	return c.Status(http.StatusOK).JSON(&fiber.Map{"status": "success", "message": "User deleted", "data": user})

}

func (uh *UserHandler) GetUserByRole(c *fiber.Ctx) error {

	role := c.Params("role")

	users, err := uh.urepo.GetUserByRole(role)

	if err != nil {

		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"status": "error", "message": "Error on getting users", "data": err})

	}

	return c.Status(http.StatusOK).JSON(&fiber.Map{"status": "success", "message": "Users found", "data": users})

}
