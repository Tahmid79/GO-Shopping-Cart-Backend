package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DBUri string = "mongodb://localhost:27017/?readPreference=primary&appname=MongoDB%20Compass&directConnection=true&ssl=false"
var DBName string = "test"
var DB *mongo.Client = ConnectDB()
var timeout time.Duration = 10 * time.Second

func ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(DBUri))
	checkErr(err)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	checkErr(err)

	err = client.Ping(ctx, nil)
	checkErr(err)

	fmt.Println("Connected to MongoDB")
	return client
}

func GetCollection(collectionName string) *mongo.Collection {
	collection := DB.Database(DBName).Collection(collectionName)
	return collection
}

func GetContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), timeout)
	return ctx
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
