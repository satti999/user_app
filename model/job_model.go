package model

import (
	"time"

	"gorm.io/gorm"
)

type Job struct {
	ID              uint           `gorm:"primaryKey" json:"jobId"`
	Title           string         `gorm:"not null" json:"title"`
	Description     string         `gorm:"not null" json:"description"`
	Requirements    string         `gorm:"not nill" json:"requirements"` // Array of strings for requirements
	Salary          int            `gorm:"not null" json:"salary"`
	ExperienceLevel int            `gorm:"not null" json:"experienceLevel"`
	Location        string         `gorm:"not null" json:"location"`
	JobType         string         `gorm:"not null" json:"jobType"`
	Position        int            `gorm:"not null" json:"position"`
	CompanyID       uint           `gorm:"not null" json:"companyId"`
	Company         Company        `gorm:"foreignKey:CompanyID" json:"company"`            // Foreign key referencing the Company model
	CreatedByID     uint           `gorm:"not null" json:"createdBy"`                      // Foreign key referencing the User model (who created the job)
	Applications    []Application  `gorm:"many2many:job_applications" json:"applications"` // Many-to-many relationship with Applications
	CreatedAt       time.Time      `json:"createdAt"`
	UpdatedAt       time.Time      `json:"updatedAt"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}
