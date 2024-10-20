package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email     string `gorm:"unique"`
	Password  string
	FirstName string
	LastName  string
	Age       int `gorm:"not null"`
}
