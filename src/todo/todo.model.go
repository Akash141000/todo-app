package todo

import (
	configurations "todoBackend/configs"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Todo struct {
	Id          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" validate:"required"`
	Description string             `json:"description"`
	Completed   bool               `json:"completed"`
}

var TodoModel *mongo.Collection

func init() {
	TodoModel = configurations.CreateCollection("todo")
}
