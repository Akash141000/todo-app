package user

import (
	"errors"
	"fmt"
	"todoBackend/db"
	helperservice "todoBackend/src/helperService"

	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user User) *mongo.InsertOneResult {
	hashedPassword := HashPassword(user)
	createdUser, _ := helperservice.InsertOne(db.UserModel, map[string]string{"email": user.Email, "password": hashedPassword})
	return createdUser
}

func LoginUser(user User) (interface{}, error) {

	//query filter
	filter := make(map[string]interface{})
	filter["email"] = user.Email

	loginUser, err := helperservice.FindOne(db.UserModel, &filter)
	if err != nil {
		return nil, err
	}

	hashedPassword := loginUser["password"].(string)
	isEqual := CompareHashedPassword(hashedPassword, user.Password)
	if isEqual {
		return loginUser, nil
	}
	return nil, errors.New("invalid credentials")
}

func CompareHashedPassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		fmt.Println("Compare Password>>", err)
		return false
	}
	return true
}

func HashPassword(user User) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	return string(hashedPassword)
}
