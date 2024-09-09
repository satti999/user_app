package repository

import (
	"fmt"

	"github.com/user_app/model"
)

type UserRepository struct {
	UserRepo *Reposiotry
}

func NewUserRepository(userrepo *Reposiotry) *UserRepository {
	return &UserRepository{
		UserRepo: userrepo,
	}
}
func (ur *UserRepository) CreateUser(user model.User, profile model.Profile) error {
	err := ur.UserRepo.DB.Create(&user).Error
	fmt.Println("User id in repo", user.ID)
	user_res, _ := ur.GetUserByEmail(user.Email, string(user.Role))
	profile.UserID = user_res.ID
	perr := ur.UserRepo.DB.Create(&profile).Error
	if err != nil || perr != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) GetUserByID(id uint) (model.User, error) {
	var user model.User
	err := ur.UserRepo.DB.Model(user).Preload("Profile").Find(&user, id)
	if err.Error != nil {
		return user, err.Error
	}
	return user, nil

}

func (ur *UserRepository) UserExists(email string, role string) bool {
	var user model.User
	err := ur.UserRepo.DB.Model(user).Where("email = ? AND role = ?", email, role).Find(&user).Error
	fmt.Println("user exists repo error", err, email, role)
	fmt.Println("user exists repo response", user.Email, user.Role)

	if email == user.Email && role == string(user.Role) {
		fmt.Println("user exists repo .................")
		return true

	}
	if err != nil {

		return false
	}

	return false

}

func (ur *UserRepository) GetUserByEmail(email string, role string) (model.User, error) {
	var user model.User
	err := ur.UserRepo.DB.Model(user).Preload("Profile").Where("email = ? AND role = ?", email, role).Find(&user)
	if err.Error != nil {
		return user, err.Error
	}
	return user, nil
}

func (ur *UserRepository) GetAllUsers() ([]model.User, error) {
	var users []model.User

	err := ur.UserRepo.DB.Model(model.User{}).Preload("Profile").Find(&users)
	if err.Error != nil {
		return users, err.Error
	}
	return users, nil
}

func (ur *UserRepository) UpdateUser(user model.User, profile model.Profile, id uint) error {
	User := model.User{}
	user.ID = id
	Profile := model.Profile{}

	err := ur.UserRepo.DB.Model(User).Where("id = ?", id).Updates(&user).Error
	perr := ur.UserRepo.DB.Model(Profile).Where("user_id", id).Updates(profile).Error

	if err != nil || perr != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) DeleteUser(user model.User) error {

	err := ur.UserRepo.DB.Model(user).Delete(user).Error
	if err != nil {
		return err
	}
	return nil
}
func (ur *UserRepository) UpdateUserRole(id uint, role string) error {

	err := ur.UserRepo.DB.Model(model.User{}).Where("id = ?", id).Update("role", role).Error

	if err != nil {
		return err
	}

	return nil

}

func (ur *UserRepository) GetUserByRole(role string) ([]model.User, error) {

	var users []model.User

	err := ur.UserRepo.DB.Model(model.User{}).Where("role = ?", role).Find(&users).Error

	if err != nil {
		return users, err
	}

	return users, nil
}
