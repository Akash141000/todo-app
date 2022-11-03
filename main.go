package main

import (
	"fmt"
	configurations "todoBackend/configs"
	"todoBackend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	port := "127.0.0.1:8080"
	ginD := gin.Default()

	//register all routes
	routes.UserRoutes(&ginD.RouterGroup)
	routes.TodoRoutes(&ginD.RouterGroup)

	ginD.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Default": "Hello from Todo application",
		})
	})
	error := ginD.Run(port)

	//connect DB
	configurations.ConnectDB()

	if error != nil {
		fmt.Println("ERROR OCCURED")
	}
}
