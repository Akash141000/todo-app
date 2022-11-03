package todo

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// /add - add new todos to db
func AddTodo(ginC *gin.Context) {
	fmt.Println("ADD TODOS")

	var todo Todo

	if err := ginC.BindJSON(&todo); err != nil {
		log.Fatal("Error Binding JSONs")
	}

	result := InsertTodo(todo)

	ginC.JSON(http.StatusAccepted, gin.H{"RESULT": result.InsertedID})

}

func GetAllTodo(ginC *gin.Context) {
	fmt.Println("GET TODOS")

	result := FindAll()

	ginC.JSON(http.StatusAccepted, gin.H{"Result": result})
}
