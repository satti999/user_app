package model

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type UseerRole string

const (
	AdminRole UseerRole = "admin"
	UserRole  UseerRole = "recruiter"
	GuestRole UseerRole = "student"
)

type Profile struct {
	gorm.Model
	Bio                string         `json:"bio"`
	Skills             pq.StringArray `gorm:"type:text[]" json:"skills"`
	Resume             string         `json:"resume"`
	ResumeOriginalName string         `json:"resumeOriginalName"`
	//CompanyID          uint     `json:"companyId"`
	ProfilePhoto string `json:"profilePhoto"`
	UserEmail    string `gorm:"not null"`
	UserID       uint   `gorm:"not null"`
}

type User struct {
	gorm.Model
	Name     string    `json:"name"`
	Role     UseerRole `json:"role"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Profile  Profile   `gorm:"foreignKey:UserID" json:"profile"`
	//CompanyID   uint      `json:"companyId"`
	Application Application `gorm:"foreignKey:UserID"`
	Company     Company     `gorm:"foreignKey:UserID"`
	Job         Job         `gorm:"foreignKey:CreatedByID"`
}
type UserReq struct {
	Name               string    `json:"name"`
	Role               UseerRole `json:"role"`
	Email              string    `json:"email"`
	Password           string    `json:"password"`
	Bio                string    `json:"bio"`
	Skills             []string  ` json:"skills"`
	Resume             string    `json:"resume"`
	ResumeOriginalName string    `json:"resumeOriginalName"`
	ProfilePhoto       string    `json:"profilePhoto"`
}
type UserResponse struct {
	Name               string    `json:"name"`
	Role               UseerRole `json:"role"`
	Email              string    `json:"email"`
	Bio                string    `json:"bio"`
	Skills             []string  ` json:"skills"`
	Resume             string    `json:"resume"`
	ResumeOriginalName string    `json:"resumeOriginalName"`
	ProfilePhoto       string    `json:"profilePhoto"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Role == "" {
		u.Role = UserRole
	}
	return nil
}

// type Role struct {
// 	gorm.Model
// 	ID          uint   `gorm:"primary_key"`
// 	Name        string `gorm:"size:50;not null;unique" json:"name"`
// 	Description string `gorm:"size:255;not null" json:"description"`
// }
