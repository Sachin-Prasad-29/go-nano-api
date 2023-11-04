package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/sachin-prasad-29/go-nano-api/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017"

const dbName = "netflix"

const colName = "watchlist"

// Important
var collection *mongo.Collection

// connect with mongoDb

// initialization method
func init() {
	// client option
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongoDb
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDb connection success")

	collection = client.Database(dbName).Collection(colName)

	//collection instance
	fmt.Println("Collection instance is ready")

}

// Mongo Helpers - file
// insert 1 record

func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted One movie in DB with Id : ", inserted.InsertedID)
}
