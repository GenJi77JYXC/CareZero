package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"not null"`
	Password string `gorm:"size:100;not null"`
	Phone    string `gorm:"unique"`
	Email    string `gorm:"size:100;unique"`
	Role     string `gorm:"size:100;default:'customer'"`
}
