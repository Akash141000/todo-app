package user

import (
	helperservice "todoBackend/src/helperService"

	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user User) *mongo.InsertOneResult {
	hashedPassword := HashPassword(user)
	createdUser, _ := helperservice.InsertOne(UserModel, user)
	return createdUser
}

func LoginUser(user User) *mongo.SingleResult {
	hashedPassword := HashPassword(user)
	loginUser := helperservice.FindById(UserModel, user.ID)
	return loginUser
}

func HashPassword(user User) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	return string(hashedPassword)
}
