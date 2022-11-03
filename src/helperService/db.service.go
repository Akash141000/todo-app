package helperservice

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindAll(model *mongo.Collection) []primitive.M {
	fmt.Println("FIND All", model.Name())
	ctx := context.Background()
	cursor, err := model.Find(ctx, bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var allDocs []bson.M

	//All will close the cursor
	err = cursor.All(context.Background(), &allDocs)
	if err != nil {
		log.Fatal("Error while getting docs from cursor")
	}

	return allDocs
}

func InsertOne(model *mongo.Collection, data interface{}) (*mongo.InsertOneResult, error) {
	fmt.Println("Insert One", model.Name())
	// if !todo.Completed {
	// 	todo.Completed = false
	// }
	insertedTodo, err := model.InsertOne(context.Background(), data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("INSERTED >>", insertedTodo.InsertedID)
	return insertedTodo, err
}

func FindById(model *mongo.Collection, id primitive.ObjectID) *mongo.SingleResult {
	fmt.Println("FIND ONE", model.Name())
	ctx := context.Background()
	// userId, err := primitive.ObjectIDFromHex(data.ID)
	foundUser := model.FindOne(ctx, bson.M{"_id": id})
	return foundUser
}
