package utils

import (
	"time"

	"gorm.io/gorm"
)

type Application struct {
	ID     uint   `json:"applicationId"`
	JobID  uint   `json:"jobId"`       // Foreign key referencing the Job model
	UserID uint   `json:"applicantId"` // Foreign key referencing the User model
	Status string `json:"status"`
}

type Job struct {
	JobID           uint           `json:"jobId"`
	Title           string         `json:"title"`
	Description     string         `json:"description"`
	Requirements    string         `json:"requirements"`
	Salary          int            `json:"salary"`
	ExperienceLevel int            `json:"experienceLevel"`
	Location        string         `json:"location"`
	JobType         string         `json:"jobType"`
	Position        int            `json:"position"`
	CreatedByID     uint           `json:"createdBy"`
	Applications    []Application  `gorm:"foreignKey:JobID" json:"applications"`
	CreatedAt       time.Time      `json:"createdAt"`
	UpdatedAt       time.Time      `json:"updatedAt"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}
