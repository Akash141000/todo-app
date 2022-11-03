package user

import (
	configurations "todoBackend/configs"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email    string             `json:"email" validate:"required"`
	Password string             `json:"password" validate:"required"`
}

var UserModel *mongo.Collection

func init() {
	UserModel = configurations.CreateCollection("user")
}
