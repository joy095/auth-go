package routes

import (
	"auth.com/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
    userController := controllers.NewUserController(db)

    userRoutes := router.Group("/user")
    {
        userRoutes.POST("/register", userController.Register)
    }

}
