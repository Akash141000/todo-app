package user

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Login(ginC *gin.Context) {
	var user User
	validate := validator.New(validator.WithRequiredStructEnabled())

	if err := ginC.BindJSON(&user); err != nil {
		ginC.JSON(http.StatusNotAcceptable, gin.H{"message": "Invalid credentials"})
		return
	}

	validationErr := validate.Struct(user)
	if validationErr != nil {
		ginC.JSON(http.StatusBadRequest, gin.H{"message": "Required fields missing!"})
		return
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
