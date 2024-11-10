package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username   		string `json:"username" binding:"required,min=3,max=20" gorm:"unique"`
	Email      		string `json:"email" binding:"required,email" gorm:"unique"`
	FirstName 		string `json:"firstName" binding:"required,firstname" `
	LastName 		string `json:"lastName" binding:"required,lastname"`
	PasswordHash   	string `json:"password" binding:"required,min=8"`
	IsVerified 		bool   `json:"isVerified"`
}


