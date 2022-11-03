package configurations

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Database {
	fmt.Println("Connecting to DB....")
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)

	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	db := client.Database("TODO")

	//create collection

	if err != nil {
		fmt.Println("Mongo error")
		log.Fatal(err)
	}
	fmt.Println("Connection to DB successfull")
	return db
}

// save db instance
var db *mongo.Database = ConnectDB()

func CreateCollection(collectionName string) *mongo.Collection {
	return db.Collection(collectionName)
}
