package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func GetBDCollection() (*mongo.Collection, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("localhost"))
	if err != nil {
		return nil, err
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	database, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	fmt.Println("basedsdedatso", database)

	collection := client.Database("Golang").Collection("User")
	fmt.Println(collection)

	list, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var abs []bson.M
	if err = list.All(ctx, &abs); err != nil {
		log.Fatal(err)
	}

	fmt.Print(abs)
	return collection, nil
}
