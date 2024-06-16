package models

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"unique;not null" json:"username" validate:"required,min=3,max=32"`
	Email     string `gorm:"unique;not null" json:"email" validate:"required,email"`
	Password  string `gorm:"not null" json:"password" validate:"required,min=6"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func MigrateUsers(db *gorm.DB) {
	err := db.AutoMigrate(&User{})
	if err != nil {
		log.Fatalf("Failed to migrate the schema: %v", err)
	}
}
