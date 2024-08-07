package model

import (
	"gorm.io/gorm"
)

type UseerRole string

const (
	AdminRole UseerRole = "admin"
	UserRole  UseerRole = "user"
	GuestRole UseerRole = "guest"
)

type User struct {
	gorm.Model
	Name string `json:"name"`
	Role UseerRole
	// RoleID   uint   `gorm:"index" json:"role_id"`
	Email string `json:"email"`
	// Role     Role   `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Password string `json:"password"`
}

// type Role struct {
// 	gorm.Model
// 	ID          uint   `gorm:"primary_key"`
// 	Name        string `gorm:"size:50;not null;unique" json:"name"`
// 	Description string `gorm:"size:255;not null" json:"description"`
// }
