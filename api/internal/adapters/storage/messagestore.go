package storage

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/lindseypoche/SELU_ACM/api/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongoRepo struct {
	client   *mongo.Client
	database string
	timeout  time.Duration
}

func newMongoClient(mongoURL string, mongoTimeout int) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(mongoTimeout)*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return client, nil
}

// NewMongoRepo creates a new mongo repository with batteries included
func NewMongoRepo(mongoURL, mongoDB string, mongoTimeout int) domain.MessageRepository {
	repo := &mongoRepo{
		timeout:  time.Duration(mongoTimeout) * time.Second,
		database: mongoDB,
	}
	client, err := newMongoClient(mongoURL, mongoTimeout)
	if err != nil {
		panic(err)
	}
	repo.client = client
	return repo
}

// Save attempts to save a blog into the database
func (repo *mongoRepo) Save(message domain.Message) (*domain.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), repo.timeout)
	defer cancel()

	collection := repo.client.Database("acm").Collection("messages")

	// no need to return message
	_, err := collection.InsertOne(ctx, &message)
	if err != nil {
		return nil, err
	}

	m := &domain.Response{Success: fmt.Sprintf("%s successfully stored in db", message.ID)}
	return m, nil
}

// GetByID attempts to get a blog by id from the database
func (repo *mongoRepo) GetByID(messageID string) (*domain.Message, error) {

	// id, err := primitive.ObjectIDFromHex(messageID)
	// if err != nil {
	// 	return nil, err
	// }
	ctx, cancel := context.WithTimeout(context.Background(), repo.timeout)
	defer cancel()

	collection := repo.client.Database("acm").Collection("messages")

	message := &domain.Message{}
	result := collection.FindOne(ctx, bson.M{"id": messageID})

	err := result.Decode(message)
	if err != nil {
		return nil, err
	}
	return message, nil
}

// GetAll gets all blogs in the database
func (repo *mongoRepo) GetAll() (*[]domain.Message, error) {

	ctx, cancel := context.WithTimeout(context.Background(), repo.timeout)
	defer cancel()

	collection := repo.client.Database("acm").Collection("messages")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var messages []domain.Message
	for cursor.Next(ctx) {
		var message domain.Message
		cursor.Decode(&message)
		messages = append(messages, message)
	}

	if len(messages) < 1 {
		return nil, errors.New("no messages found")
	}
	return &messages, nil
}

// GetByAuthor ...
func (repo *mongoRepo) GetByAuthor() error {
	return nil
}

// Update ...
func (repo *mongoRepo) Update() error {
	return nil
}

// Delete ...
func (repo *mongoRepo) Delete() error {
	return nil
}
