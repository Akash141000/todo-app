package main

import (
	"fmt"
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
	routes.TodoRoutes(&ginD.RouterGroup)

	ginD.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Default": "Hello from Todo application",
		})
	})
	error := ginD.Run(port)

	if error != nil {
		fmt.Println("ERROR OCCURED")
	}
}
