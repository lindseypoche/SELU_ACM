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
	inDevelopment      = true
	localDevHost       = "mongodb://127.0.0.1:27017"
	localContainerHost = "mongodb://0.0.0.0:27017"
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

	var uri string
	if inDevelopment {
		uri = localDevHost
	} else {

		uri := os.Getenv("MONGO_HOST")
		if uri == "" {
			uri = localContainerHost
		}
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

	var clientOptions *options.ClientOptions
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if !inDevelopment {

		clientOptions = options.Client().ApplyURI(uri).
			SetAuth(options.Credential{
				// AuthSource: "admin",
				Username: os.Getenv("MONGO_INITDB_ROOT_USERNAME"),
				Password: os.Getenv("MONGO_INITDB_ROOT_PASSWORD"),
			})

	} else {
		clientOptions = options.Client().ApplyURI(uri)
	}

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
