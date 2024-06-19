package domain

import (
	"gorm.io/gorm"
)

// User represents a user in the database
type User struct {
	gorm.Model
	Id       int    `json:"id" gorm:"primaryKey"`
	Name     string `gorm:" not null" json:"name"`
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
}
