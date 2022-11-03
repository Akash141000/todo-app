package user

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(ginC *gin.Context) {
	fmt.Println("LOGIN")
	var user User
	if err := ginC.BindJSON(&user); err != nil {
		log.Fatal("Error Binding JSONs")
	}

	result := LoginUser(user)

	ginC.JSON(http.StatusAccepted, gin.H{"RESULT": result})
}

func Signup(ginC *gin.Context) {
	fmt.Println("SIGNUP")
	var user User
	if err := ginC.BindJSON(&user); err != nil {
		log.Fatal("Error Binding JSONs")
	}

	result := CreateUser(user)

	ginC.JSON(http.StatusAccepted, gin.H{"RESULT": result.InsertedID})

}
