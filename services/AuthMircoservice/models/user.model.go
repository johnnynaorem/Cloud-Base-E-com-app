package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Password string
	Email    string `gorm:"unique"`
	Phone    string `gorm:"unique"`
}
