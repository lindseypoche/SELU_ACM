package mongo

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	defaultHost = "mongodb://0.0.0.0:27017"
	// defaultHost = `mongodb://acm-mongo:27017`
)

var (
	client *mongo.Client
	err    error
)

func GetClient() *mongo.Client {

	if client != nil {
		return client
	}

	uri := os.Getenv("MONGO_HOST")
	if uri == "" {
		uri = defaultHost
	}

	// h, err := url.Parse(uri)
	// if err != nil {
	// 	panic(err)
	// }

	// // check scheme
	// if h.Scheme != "mongodb" {
	// 	log.Println("error: h.Scheme =", h.Scheme)
	// 	panic(err)
	// }

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI(uri).
		SetAuth(options.Credential{
			// AuthSource: "admin",
			Username: os.Getenv("MONGO_INITDB_ROOT_USERNAME"),
			Password: os.Getenv("MONGO_INITDB_ROOT_PASSWORD"),
		})

	defer cancel()

	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("DB connection error: ", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Ping err: ", err)
	}

	fmt.Println("Connected to mongodb")
	return client
}
