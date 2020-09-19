package models

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var Connection *mongo.Client
var ctx = context.TODO()

func GetConnection() *mongo.Client {
	if Connection != nil {
		return Connection
	}

	clientOptions := options.Client().ApplyURI("mongodb://db:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	Connection = client

	return Connection
}
