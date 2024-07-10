package main

import (
	"log"
	db "todoBackend/db"
	"todoBackend/middleware"
	"todoBackend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	port := "127.0.0.1:8080"
	ginD := gin.Default()

	//connect DB
	db.ConnectDB()

	//register all routes
	routes.UserRoutes(&ginD.RouterGroup)

	//Auth middleware for all below routes
	ginD.Use(middleware.AuthGuard)

	//guarded routes
	routes.TodoRoutes(&ginD.RouterGroup)

	//testing route
	ginD.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Default": "Hello from Todo application",
		})
	})

	//start the server
	serverError := ginD.Run(port)

	if serverError != nil {
		log.Panic("unable to start the server", serverError)
	}
}
