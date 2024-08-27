package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/user_app/config"
	"github.com/user_app/middleware"
	"github.com/user_app/model"
	"github.com/user_app/repository"
	"github.com/user_app/utils"
)

const (
	oauthGoogleUserInfoURL = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="
	// oauthGithubUserURL       = "https://api.github.com/user"
	// oauthGithubUserEmailsURL = "https://api.github.com/user/emails"
)

type UserHandlerInterface interface {
	CreateUser(c *fiber.Ctx) error
	LoginHandler(c *fiber.Ctx) error
	GetUserByID(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
	GetAllUsers(c *fiber.Ctx) error
	GetUserByRole(c *fiber.Ctx) error
	GetUserByEmail(c *fiber.Ctx) error
	GoogleSignin(c *fiber.Ctx) error
	GoogleCallback(c *fiber.Ctx) error
}

type UserHandler struct {
	urepo *repository.UserRepository
}

func NewUserHandler(useRrepo *repository.UserRepository) UserHandlerInterface {
	return &UserHandler{
		urepo: useRrepo,
	}
}

func (uh *UserHandler) CreateUser(c *fiber.Ctx) error {
	imagUrl := utils.GetProfileUrl()
	User := model.User{}
	Profile := model.Profile{}
	UserReq := model.UserReq{}

	err := c.BodyParser(&UserReq)
	fmt.Println("user", User.Name)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{"status": "error", "message": "Error on request", "data": err})

	}

	exists := uh.urepo.UserExists(User.Email)

	if exists {
		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "User already exists"})
	}
	User.Name = UserReq.Name
	User.Email = UserReq.Email
	User.Role = UserReq.Role
	User.Password = UserReq.Password
	Profile.UserEmail = UserReq.Email
	Profile.Bio = UserReq.Bio
	Profile.Skills = UserReq.Skills
	Profile.Resume = UserReq.Resume
	Profile.ResumeOriginalName = UserReq.ResumeOriginalName
	Profile.ProfilePhoto = imagUrl
	pass := User.Password
	hashedPass := utils.HashAndSalt(pass)
	User.Password = hashedPass

	err = uh.urepo.CreateUser(User, Profile)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Error on creating user"})

	}

	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "User has been created successfully",
		"status":  "success"})

}
func (uh *UserHandler) LoginHandler(c *fiber.Ctx) error {

	var u model.User

	err := c.BodyParser(&u)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Error on login request"})

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

	if email == "" {

		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"status": "error", "message": "Email is required", "data": nil})
	}
	fmt.Println("email", email)

	user, err := uh.urepo.GetUserByEmail(email)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"status": "error", "message": "Error on getting user", "data": err})
	}

	return c.Status(http.StatusOK).JSON(&fiber.Map{"status": "success", "message": "User found", "data": user})

}

func (uh *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	UserRes := model.UserResponse{}
	users, err := uh.urepo.GetAllUsers()

	if err != nil {

		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"status": "error", "message": "Error on getting users", "data": err})
	}
	UserRes.Name = users[0].Name
	UserRes.Email = users[0].Email
	UserRes.Role = users[0].Role
	UserRes.Bio = users[0].Profile.Bio
	UserRes.Skills = users[0].Profile.Skills
	UserRes.Resume = users[0].Profile.Resume
	UserRes.ResumeOriginalName = users[0].Profile.ResumeOriginalName
	UserRes.ProfilePhoto = users[0].Profile.ProfilePhoto

	return c.Status(http.StatusOK).JSON(&fiber.Map{"status": "success", "message": "Users found", "profile": UserRes})

}

func (uh *UserHandler) UpdateUser(c *fiber.Ctx) error {

	userID, err := c.ParamsInt("id")
	imagUrl := utils.GetProfileUrl()
	resumeUrl := utils.GetResumeUrl()
	User := model.User{}
	Profile := model.Profile{}
	UserReq := model.UserReq{}

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{"status": "error", "message": "Id is invalid", "data": err})
	}

	user_ID := uint(userID)

	if err := c.BodyParser(&UserReq); err != nil {

		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"status": "error", "message": "Error on request", "data": err})

	}
	fmt.Println("User skills are", UserReq.Skills)
	fmt.Printf("%T\n", UserReq.Skills)
	User.Name = UserReq.Name
	User.Email = UserReq.Email
	User.Password = UserReq.Password
	Profile.Bio = UserReq.Bio
	Profile.Skills = UserReq.Skills
	Profile.Resume = resumeUrl
	Profile.UserID = user_ID
	Profile.ResumeOriginalName = UserReq.ResumeOriginalName
	Profile.ProfilePhoto = imagUrl

	if User.Password != "" {
		pass := utils.HashAndSalt(User.Password)

		User.Password = pass
	}

	err = uh.urepo.UpdateUser(User, Profile, user_ID)

	if err != nil {

		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"status": "error", "message": "Error on updating user", "data": err})

	}
	User.ID = user_ID

	return c.Status(http.StatusOK).JSON(&fiber.Map{"status": "success", "message": "User updated"})

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

func (uh *UserHandler) GoogleSignin(c *fiber.Ctx) error {

	googleOAuthConfig := config.AppConfig.GoogleLoginConfig
	url := googleOAuthConfig.AuthCodeURL("randomstate")

	c.Status(fiber.StatusSeeOther)
	c.Redirect(url)
	return nil

}

func (uh *UserHandler) GoogleCallback(c *fiber.Ctx) error {

	code := c.Query("code")

	googleOAuthConfig := config.AppConfig.GoogleLoginConfig

	token, err := googleOAuthConfig.Exchange(context.Background(), code)

	if err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(
			&fiber.Map{"status": "error", "message": "Failed to login", "data": err})

	}

	response, err := http.Get(oauthGoogleUserInfoURL + token.AccessToken)

	if err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(
			&fiber.Map{"status": "error", "message": "Failed to login", "data": err})

	}
	defer response.Body.Close()
	var userInfo map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&userInfo)

	if err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(
			&fiber.Map{"status": "error", "message": "Failed to login", "data": err})

	}
	fmt.Println("User info", userInfo)
	email := userInfo["email"].(string)

	user, err := uh.urepo.GetUserByEmail(email)

	if err != nil || user.Email == "" {

		return c.Status(fiber.StatusInternalServerError).JSON(
			&fiber.Map{"status": "error", "message": "Please sign up first", "data": err})

	} else {

		userToken, erre := middleware.CreateToken(user)

		if erre != nil {

			return c.Status(fiber.StatusInternalServerError).JSON(
				&fiber.Map{"status": "error", "message": "Failed to login", "data": err})

		}
		cookie := fiber.Cookie{
			Name:     "jwt",
			Value:    userToken,
			Expires:  time.Now().Add(time.Hour * 24),
			HTTPOnly: true,
		}

		c.Cookie(&cookie)

		return c.JSON(fiber.Map{"status": "success",
			"message": "User logged in successfully",
			"data":    userToken,
			"userId":  user.ID,
			"role":    user.Role,
			"email":   user.Email})

	}

}
