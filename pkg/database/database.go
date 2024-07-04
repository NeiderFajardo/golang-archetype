package database

import (
	"context"
	"log"
	"time"

	"github.com/NeiderFajardo/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDatabase struct {
	Collection *mongo.Collection
	Client     *mongo.Client
	Ctx        context.Context
	Cancel     context.CancelFunc
}

func NewMongoClient(config *config.MongoConfig) *MongoDatabase {
	// Initialize the database
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(config.URI()))
	if err != nil {
		log.Fatalf("connection error :%v", err)
	}
	collection := mongoClient.Database(config.Database()).Collection(config.Collection())

	return &MongoDatabase{
		Collection: collection,
		Client:     mongoClient,
		Ctx:        ctx,
		Cancel:     cancel,
	}
}

func CloseConnection(db *MongoDatabase) {
	defer func() {
		db.Cancel()
		if err := db.Client.Disconnect(db.Ctx); err != nil {
			log.Fatalf("mongodb disconnect error : %v", err)
		}
	}()
}
