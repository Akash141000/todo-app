package user

import (
	helperservice "todoBackend/src/helperService"

	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(user User) *mongo.InsertOneResult {
	createdUser, _ := helperservice.InsertOne(UserModel, user)
	return createdUser
}

func LoginUser(user User) *mongo.SingleResult {
	loginUser := helperservice.FindById(UserModel, user.ID)
	return loginUser
}
