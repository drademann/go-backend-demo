package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var client *mongo.Client
var database *mongo.Database

func Open(dbName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return fmt.Errorf("unable to connect to mongo: %w", err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		return fmt.Errorf("unable to ping mongo: %w", err)
	}
	database = client.Database(dbName)
	return nil
}

func Close() {
	err := client.Disconnect(context.Background())
	if err != nil {
		log.Fatalf("failed to disconnect from mongo: %v", err)
	}
}

func Collection(name string) *mongo.Collection {
	return database.Collection(name)
}

func Find(collection *mongo.Collection, result interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)
	return cursor.All(ctx, result)
}
