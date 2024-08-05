package repository

import "gorm.io/gorm"

type Reposiotry struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Reposiotry {
	return &Reposiotry{
		DB: db,
	}
}
