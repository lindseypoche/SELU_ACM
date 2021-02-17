package storage

import (
	"context"
	"errors"
	"time"

	"github.com/lindseypoche/SELU_ACM/api/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
func NewMongoRepo(mongoURL, mongoDB string, mongoTimeout int) domain.BlogRepository {
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
func (repo *mongoRepo) Save(blog domain.Blog) (map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), repo.timeout)
	defer cancel()

	collection := repo.client.Database("acm").Collection("blogs")

	result, err := collection.InsertOne(ctx, &blog)
	if err != nil {
		return nil, err
	}

	r := result.InsertedID.(primitive.ObjectID).Hex()
	m := make(map[string]interface{})
	m["id"] = r
	return m, nil
}

// GetByID attempts to get a blog by id from the database
func (repo *mongoRepo) GetByID(blogID string) (*domain.Blog, error) {

	id, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), repo.timeout)
	defer cancel()

	collection := repo.client.Database("acm").Collection("blogs")

	blog := &domain.Blog{}
	result := collection.FindOne(ctx, bson.M{"_id": id})

	err = result.Decode(blog)
	if err != nil {
		return nil, err
	}
	return blog, nil
}

// GetAll gets all blogs in the database
func (repo *mongoRepo) GetAll() (*[]domain.Blog, error) {

	ctx, cancel := context.WithTimeout(context.Background(), repo.timeout)
	defer cancel()

	collection := repo.client.Database("acm").Collection("blogs")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var blogs []domain.Blog
	for cursor.Next(ctx) {
		var blog domain.Blog
		cursor.Decode(&blog)
		blogs = append(blogs, blog)
	}

	if len(blogs) < 1 {
		return nil, errors.New("no blogs found")
	}
	return &blogs, nil
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
