package user

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(ginC *gin.Context) {
	fmt.Println("API >> LOGIN")
	var user User
	if err := ginC.BindJSON(&user); err != nil {
		log.Fatal("Error Binding JSONs")
	}

	userFound, err := LoginUser(user)
	if err != nil {
		ginC.JSON(http.StatusNotAcceptable, gin.H{"message": "Invalid credentials"})
		return
	}
	ginC.JSON(http.StatusAccepted, gin.H{"message": "User logged in Successfully", "data": userFound})
}

func Signup(ginC *gin.Context) {
	fmt.Println("API >> SIGNUP")
	var user User
	if err := ginC.BindJSON(&user); err != nil {
		log.Fatal("Error Binding JSONs")
	}

	createdUser := CreateUser(user)
	if createdUser == nil {
		ginC.JSON(http.StatusNotAcceptable, gin.H{"message": "Unable to create User"})
		return
	}
	ginC.JSON(http.StatusAccepted, gin.H{"message": "User created successfully"})

}
