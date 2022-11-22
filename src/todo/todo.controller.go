package todo

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// /add - add new todos to db
func AddTodoHandler(ginC *gin.Context) {

	var todo Todo

	if err := ginC.BindJSON(&todo); err != nil {
		log.Fatal("Error Binding JSONs")
	}

	fmt.Println("INSERT TODO>>", todo)

	result := InsertTodo(todo)

	ginC.JSON(http.StatusAccepted, gin.H{"RESULT": result.InsertedID})

}

func GetAllTodoHandler(ginC *gin.Context) {
	fmt.Println("GET TODOS")

	result := FindAll()

	ginC.JSON(http.StatusAccepted, gin.H{"todoList": result, "message": "Todo list"})
}

func UpdateTodoHandler(ginC *gin.Context) {
	var todoData struct {
		Id   string
		Data interface {
		}
	}

	if err := ginC.BindJSON(&todoData); err != nil {
		log.Fatal("ERROR", err)
	}
	fmt.Println("UPDATE TODO>>", todoData)

	// var todo Todo

	if doc, err := UpdateTodo(todoData.Id, todoData.Data); err != nil {
		fmt.Println("DOC", doc)
		log.Fatal(err)
	}

	ginC.JSON(http.StatusAccepted, gin.H{"RESULT": "Doc updated"})
}
