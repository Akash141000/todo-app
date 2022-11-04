package routes

import (
	"todoBackend/src/todo"

	"github.com/gin-gonic/gin"
)

func TodoRoutes(c *gin.RouterGroup) {
	c.GET("/getTodos", todo.GetAllTodoHandler)
	c.POST("/addTodo", todo.AddTodoHandler)
	c.POST("/updateTodo", todo.UpdateTodoHandler)
}
