package user

import (
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

func LoginUser(user User) (*struct{}, error) {
	hashedPassword := HashPassword(user)

	loginUser, err := helperservice.FindOne(db.UserModel, map[string]string{"email": user.Email, "password": hashedPassword})
	return loginUser, err
}

func HashPassword(user User) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	return string(hashedPassword)
}
