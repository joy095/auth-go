package controllers

import (
	"net/http"

	"auth.com/models"
	"auth.com/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
    DB *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
    return &UserController{DB: db}
}

// Register handles user registration
func (uc *UserController) Register(c *gin.Context) {
	var request struct {
		Username string 	`json:"username" binding:"required"`
		Password string 	`json:"password" binding:"required"`
		FirstName string 	`json:"firstname" binding:"required"`
		LastName string 	`json:"lastname" binding:"required"`
		Email string 		`json:"email" binding:"required"`
	}

    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    passwordHash, err := utils.HashPassword(request.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
        return
    }

    user := models.User{
		Username: request.Username, 
		PasswordHash: passwordHash, 
		Email: request.Email, 
		FirstName: request.FirstName, 
		LastName: request.LastName,
	}
    if err := uc.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}
