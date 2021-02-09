package blogs

import (
	"context"
	"errors"
	"time"

	"github.com/lindseypoche/SELU_ACM/api/clients/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Save attempts to save a blog into the database
func (blog *Blog) Save() (map[string]interface{}, error) {

	collection := mongo.Client.GetDatabaseAndCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

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
func (blog *Blog) GetByID(blogID string) error {

	id, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return err
	}

	collection := mongo.Client.GetDatabaseAndCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result := collection.FindOne(ctx, bson.M{"_id": id})

	err = result.Decode(blog)
	if err != nil {
		return err
	}
	return nil
}

// GetAll gets all blogs in the database
func GetAll() (*[]Blog, error) {

	collection := mongo.Client.GetDatabaseAndCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var blogs []Blog
	for cursor.Next(ctx) {
		var blog Blog
		cursor.Decode(&blog)
		blogs = append(blogs, blog)
	}

	if len(blogs) < 1 {
		return nil, errors.New("no blogs found")
	}
	return &blogs, nil
}

// GetByAuthor ...
func (blog *Blog) GetByAuthor() error {
	return nil
}

// Update ...
func (blog *Blog) Update() error {
	return nil
}

// Delete ...
func (blog *Blog) Delete() error {
	return nil
}
