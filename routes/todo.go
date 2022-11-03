package routes

import (
	"todoBackend/src/todo"

	"github.com/gin-gonic/gin"
)

func TodoRoutes(c *gin.RouterGroup) {
	c.GET("/getTodos", todo.GetAllTodo)
	c.POST("/addTodo", todo.AddTodo)
}
