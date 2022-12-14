package todo

import (
	"context"
	"fmt"
	"log"
	db "todoBackend/db"
	helperservice "todoBackend/src/helperService"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindAll() []primitive.M {
	fmt.Printf("FIND All")
	ctx := context.Background()
	cursor, err := db.TodoModel.Find(ctx, bson.D{{}})
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

func InsertTodo(todo Todo) *mongo.InsertOneResult {
	fmt.Println("Insert One")
	if !todo.Completed {
		todo.Completed = false
	}
	insertedTodo, err := helperservice.InsertOne(db.TodoModel, todo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("INSERTED >>", insertedTodo.InsertedID)
	return insertedTodo
}

func FindOne() {
	fmt.Println("FIND ONE")

}

func UpdateTodo(todoId string, data interface{}) (map[string]interface{}, error) {
	id, _ := primitive.ObjectIDFromHex(todoId)

	filter := bson.M{"_id": id}
	update := bson.M{"$set": data}
	updatedDoc, err := helperservice.UpdateOne(db.TodoModel, filter, update)

	return updatedDoc, err
}
