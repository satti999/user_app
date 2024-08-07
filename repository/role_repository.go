package repository

// import (
// 	"github.com/user_app/model"
// )

// type RoleRepository struct {
// 	RoleRepo *Reposiotry
// }

// func NewRoleRepository(rolerepo *Reposiotry) *RoleRepository {
// 	return &RoleRepository{
// 		RoleRepo: rolerepo,
// 	}
// }

// func (r *RoleRepository) CreateRole(Role *model.Role) (err error) {
// 	err = r.RoleRepo.DB.Create(Role).Error
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// // Get all roles
// func (r *RoleRepository) GetRoles(Role *[]model.Role) (err error) {
// 	err = r.RoleRepo.DB.Find(Role).Error
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// // Get role by id
// func (r *RoleRepository) GetRole(id int) (rol model.Role, err error) {
// 	role := model.Role{}
// 	err = r.RoleRepo.DB.Where("id = ?", id).First(&role).Error
// 	if err != nil {
// 		return model.Role{}, err
// 	}
// 	return role, err
// }

// // Update role
// func (r *RoleRepository) UpdateRole(Role *model.Role, id uint) (err error) {

// 	err = r.RoleRepo.DB.Model(model.Role{}).Where("id = ?", id).Updates(Role).Error
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
