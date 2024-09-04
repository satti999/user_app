package model

import (
	"time"

	"gorm.io/gorm"
)

type Company struct {
	ID          uint   `gorm:"primaryKey" json:"companyId"`
	Name        string `gorm:"not null;unique" json:"name"`
	Description string `json:"description"`
	Website     string `json:"website"`
	Location    string `json:"location"`
	Logo        string `json:"logo"`
	UserID      uint   `gorm:"not null" json:"userId"` // Foreign key referencing the User model
	//ProfileID   uint           `gorm:"not null" json:"profileId"` // Foreign key referencing the User model
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type CompaniesResponse struct {
	Companies []Company `json:"companies"`
}
