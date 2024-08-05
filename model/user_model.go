package model

import (
	"gorm.io/gorm"
)

const (
	RoleAdmin = "admin"
	RoleUser  = "user"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Password string `json:"password"`
}

func isAdmin(user User) bool {
	return user.Role == RoleAdmin
}
