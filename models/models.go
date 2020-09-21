package models

import (
	"gorm.io/gorm"
)

// User represents the core user model
type User struct {
	gorm.Model
	Email        string `gorm:"primaryKey"`
	Password     string
	PasswordSalt string
}
