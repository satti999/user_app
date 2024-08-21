package databasego

import (
	"fmt"

	"github.com/user_app/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	SSLMode  string
}

func ConnectDB(config *Config) *gorm.DB {
	// dsn := "host=localhost user=postgres password=admin dbname=user_app port=5432 sslmode=disable"
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connected successfully")

	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&model.User{},
		&model.Profile{},
		&model.Company{},
		&model.Application{},
		&model.Job{})
	fmt.Println("Database migrated successfully")

}
