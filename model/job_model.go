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
	Company         Company        `gorm:"foreignKey:CompanyID" json:"company"`  // Foreign key referencing the Company model
	CreatedByID     uint           `gorm:"not null" json:"createdBy"`            // Foreign key referencing the User model (who created the job)
	Applications    []Application  `gorm:"foreignKey:JobID" json:"applications"` // Many-to-many relationship with Applications
	CreatedAt       time.Time      `json:"createdAt"`
	UpdatedAt       time.Time      `json:"updatedAt"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}
type ApplicationData struct {
	ID     uint   `json:"applicationId"`
	JobID  uint   `json:"jobId"`       // Foreign key referencing the Job model
	UserID uint   `json:"applicantId"` // Foreign key referencing the User model
	Status string `json:"status"`
}

type JobWithApplications struct {
	JobID           uint              `json:"jobId"`
	Title           string            `json:"title"`
	Description     string            `json:"description"`
	Requirements    string            `json:"requirements"`
	Salary          int               `json:"salary"`
	ExperienceLevel int               `json:"experienceLevel"`
	Location        string            `json:"location"`
	JobType         string            `json:"jobType"`
	Position        int               `json:"position"`
	CreatedByID     uint              `json:"createdBy"`
	Applications    []ApplicationData `gorm:"foreignKey:JobID" json:"applications"`
	CreatedAt       time.Time         `json:"createdAt"`
	UpdatedAt       time.Time         `json:"updatedAt"`
	DeletedAt       gorm.DeletedAt    `gorm:"index" json:"-"`
}
