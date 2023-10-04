package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nguyen997/gin-gorm-rest/config"
	"github.com/nguyen997/gin-gorm-rest/routes"
)

//	func init() {
//		initializers.SyncDatabase()
//	}
func main() {
	route := gin.Default()
	config.Connect()
	routes.UserRoute(route)
	route.Run(":8080")
}
