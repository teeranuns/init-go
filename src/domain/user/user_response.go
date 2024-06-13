package domain

import (
	"gorm.io/gorm"
)

// User represents a user in the database
type User struct {
	gorm.Model
	Name     string `gorm:"size:100;not null" json:"name"`
	Email    string `gorm:"size:100;unique;not null" json:"email"`
	Password string `gorm:"size:100;not null" json:"password"`
}
