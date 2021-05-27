package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetBDCollection(collectionName string) (*mongo.Collection, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://tezca:mongo@jwt.mmjrl.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	collection := client.Database("Golang").Collection(collectionName)
	return collection, nil
}
