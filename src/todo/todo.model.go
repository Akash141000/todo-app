package todo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	Id          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" validate:"required"`
	Description string             `json:"description"`
	Completed   bool               `json:"completed"`
}
