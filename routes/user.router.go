package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nguyen997/gin-gorm-rest/controller"
	"github.com/nguyen997/gin-gorm-rest/middleware"
)

func UserRoute(router *gin.Engine) {
	router.GET("/", controller.GetUser)
	router.POST("/", controller.Signup)
	router.POST("/login", controller.Login)
	router.GET("/validate", middleware.RequireAuth, controller.Validate)
	router.PUT("/:id", controller.UpdateUser)
	router.DELETE("/:id", controller.DeleteUser)
}
