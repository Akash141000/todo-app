package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// save db instance
var db *mongo.Database
var UserModel *mongo.Collection
var TodoModel *mongo.Collection

func ConnectDB() *mongo.Database {
	fmt.Println("Connecting to DB....")
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	//create mongodb client
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println("MONGODB ERROR ")
		log.Panic(err)
		panic(err)
	}

	fmt.Println("Checking Mongodb connection")

	mongoErr := client.Ping(ctx, nil)

	if mongoErr != nil {
		log.Fatal("MONGODB NOT RUNNING!")
		panic(mongoErr)
	}
	db = client.Database("TODO")
	fmt.Println("Connection to DB successfull")
	//create mongo collections
	UserModel = CreateCollection("user")
	TodoModel = CreateCollection("todo")
	return db
}

func CreateCollection(collectionName string) *mongo.Collection {
	if db == nil {
		log.Fatal("Database not connected!")
	}
	return db.Collection(collectionName)
}
