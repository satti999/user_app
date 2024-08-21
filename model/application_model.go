package model

import (
	"time"

	"gorm.io/gorm"
)

type ApplicationStatus string

const (
	Pending  ApplicationStatus = "pending"
	Accepted ApplicationStatus = "accepted"
	Rejected ApplicationStatus = "rejected"
)

type Application struct {
	ID        uint              `gorm:"primaryKey" json:"applicationId"`
	JobID     uint              `gorm:"not null" json:"jobId"`       // Foreign key referencing the Job model
	UserID    uint              `gorm:"not null" json:"applicantId"` // Foreign key referencing the User model
	Status    ApplicationStatus `gorm:"default:pending" json:"status"`
	CreatedAt time.Time         `json:"createdAt"`
	UpdatedAt time.Time         `json:"updatedAt"`
	DeletedAt gorm.DeletedAt    `gorm:"index" json:"-"`
}
