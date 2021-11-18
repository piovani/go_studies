package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var databaseName = "go_mongo"
var collectionName = "teste"

type Task struct {
	ID   int64  `bson:"_id"`
	Text string `bson:"text"`
}

func main() {
	client := connection()

	// REGISTROS
	// res := insertRecord(client)
	res := updateRecord(client)

	if !res {
		fmt.Println("DEU RUIM")
		return
	}

	fmt.Println("AI SIMMM")
}

func connection() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://go_mongo-mongo/")

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

// REGISTROS
func insertRecord(client *mongo.Client) bool {
	var collection *mongo.Collection
	var task Task

	task.ID = 3
	task.Text = "Teste de texto3"

	collection = client.Database(databaseName).Collection(collectionName)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, task)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func updateRecord(client *mongo.Client) bool {

}
