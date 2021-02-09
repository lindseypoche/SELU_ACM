package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	// Client is the mongo client.
	Client mongoClientInterface = &mongoClient{}
)

type mongoClientInterface interface {
	setClient(*mongo.Client)
	GetDatabaseAndCollection() *mongo.Collection
}

type mongoClient struct {
	client *mongo.Client
}

// Init initializes the mongo database
func Init() {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	Client.setClient(client)
}

func (c *mongoClient) setClient(client *mongo.Client) {
	c.client = client
}

// GetDatabaseAndCollection creates the database and collection for blogs
func (c *mongoClient) GetDatabaseAndCollection() *mongo.Collection {
	return c.client.Database("acm-blogs").Collection("blogs")
}
