package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://admin:TwUFKer9YwiLQv88@cluster0.6zeco.mongodb.net/"
const dbName = "netflix"
const collectionName = "watchlist"

// most imp
var collections *mongo.Collection

// connect with DB
func init() {
	clientOption := options.Client().ApplyURI(connectionString)
	// connect to MDB
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB connected!!")
	collections = client.Database(dbName).Collection(collectionName)
	fmt.Println("collection instance ready!!")
}
