package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var databaseName = "go_mongo"
var collectionName = "teste"

type Task struct {
	ID   int64  `bson:"id"`
	Text string `bson:"text"`
}

func main() {
	// DOCUMENTS
	// res := insertRecord()
	// res := updateRecord()
	// res := deleteRecord()
	// res := listRecords()

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

func getCollectionAndCtx() (*mongo.Collection, context.Context) {
	var collection *mongo.Collection
	client := connection()

	collection = client.Database(databaseName).Collection(collectionName)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	return collection, ctx
}

// DOCUMENTS
func insertRecord() bool {
	var task Task
	task.ID = 7
	task.Text = "Teste de texto7"

	collection, ctx := getCollectionAndCtx()

	_, err := collection.InsertOne(ctx, task)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func updateRecord() bool {
	id := int64(7)
	novoTexto := "O Texto foi alterado2"

	collection, ctx := getCollectionAndCtx()

	filter := bson.M{"_id": bson.M{"$eq": id}}
	update := bson.M{"$set": bson.M{"text": novoTexto}}

	result, err := collection.UpdateOne(
		ctx,
		filter,
		update,
	)

	if err != nil {
		println(err)
		return false
	}

	println("Quantidade de registros que bateram com o filtro", result.MatchedCount)
	println("Quantidade de registros alterados", result.ModifiedCount)
	println("Quantidade de registros alterados", result.UpsertedCount)

	return true
}

func deleteRecord() bool {
	id := int64(6)
	collection, ctx := getCollectionAndCtx()

	filter := bson.M{"_id": bson.M{"$eq": id}}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		fmt.Println(err)
		return false
	}

	println("Quantidade de registros deletados", result.DeletedCount)

	return true
}

func listRecords() bool {
	var tasks []*Task

	collection, ctx := getCollectionAndCtx()

	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		fmt.Println(err)
		return false
	}

	for cur.Next(ctx) {
		var t Task
		err := cur.Decode(&t)
		if err != nil {
			fmt.Println(err)
			return false
		}

		tasks = append(tasks, &t)
	}

	fmt.Println(tasks)

	return true
}
