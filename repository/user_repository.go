package repository

import (
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
func (ur *UserRepository) CreateUser(user model.User) error {
	//if user already exists then don't create the user

	// _, err := ur.GetUserByEmail(user.Email)
	// if err nil {
	// 	return fmt.Errorf("user already exists")
	// }
	// fmt.Println("user already exists", user)
	err := ur.UserRepo.DB.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) GetUserByID(id uint) (model.User, error) {
	var user model.User
	err := ur.UserRepo.DB.Model(user).Find(&user, id)
	if err.Error != nil {
		return user, err.Error
	}
	return user, nil

}

func (ur *UserRepository) UserExists(email string) bool {
	var user model.User
	err := ur.UserRepo.DB.Model(user).Where("email = ?", email).Find(&user).Error
	if err != nil {
		return false
	}
	if user.ID == 0 {
		return false
	}
	return true
}

func (ur *UserRepository) GetUserByEmail(email string) (model.User, error) {
	var user model.User
	err := ur.UserRepo.DB.Model(user).Where("email = ?", email).Find(&user)
	if err.Error != nil {
		return user, err.Error
	}
	return user, nil
}

func (ur *UserRepository) GetAllUsers() ([]model.User, error) {
	var users []model.User
	err := ur.UserRepo.DB.Model(model.User{}).Find(&users)
	if err.Error != nil {
		return users, err.Error
	}
	return users, nil
}

func (ur *UserRepository) UpdateUser(user model.User, id uint) error {
	User := model.User{}
	user.ID = id
	err := ur.UserRepo.DB.Model(User).Where("id = ?", id).Updates(user).Error
	if err != nil {
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

func (ur *UserRepository) GetUserByRole(role string) ([]model.User, error) {

	var users []model.User

	err := ur.UserRepo.DB.Model(model.User{}).Where("role = ?", role).Find(&users).Error

	if err != nil {
		return users, err
	}

	return users, nil
}
