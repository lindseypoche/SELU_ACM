package storage

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/lindseypoche/SELU_ACM/api/internal/domain"
	"github.com/lindseypoche/SELU_ACM/api/internal/utils/errors/rest"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// const (
// lookupByAuthorID = bson.D{{"$lookup", bson.D{{}}}}
// )

type mongoRepo struct {
	client   *mongo.Client
	database string
	timeout  time.Duration
}

// newMongoAtlasClient connects to the remote mongo atlas cluster
func newMongoAtlasClient() (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://QuantaCake01:Mo26oDB.Co3!@cluster0.mis3d.mongodb.net/acm?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	return client, nil
	// defer client.Disconnect(ctx)
}

// newMongoClient connects to your local mongo client
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
func (repo *mongoRepo) GetAll() (*[]domain.Message, rest.Err) {

	ctx, cancel := context.WithTimeout(context.Background(), repo.timeout)
	defer cancel()

	collection := repo.client.Database("acm").Collection("messages")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, rest.NewInternalServerError("error initializing cursor", err)
	}
	defer cursor.Close(ctx)

	var messages []domain.Message
	for cursor.Next(ctx) {
		var message domain.Message
		cursor.Decode(&message)
		messages = append(messages, message)
	}

	// check if there are any errors with cursor
	if err := cursor.Err(); err != nil {
		restErr := rest.NewInternalServerError("error when searching database", err)
		return nil, restErr
	}

	if len(messages) < 1 {
		return nil, rest.NewNotFoundError("can not find any messages in database")
	}
	return &messages, nil
}

// GetByAuthor gets all the posts created by an author
func (repo *mongoRepo) GetByAuthor(authorID string) (*[]domain.Message, rest.Err) {
	ctx, cancel := context.WithTimeout(context.Background(), repo.timeout)
	defer cancel()

	collection := repo.client.Database("acm").Collection("messages")

	cursor, err := collection.Find(ctx, bson.M{
		"author.id": authorID,
	})
	defer cursor.Close(ctx)

	messages := []domain.Message{}

	// loop through cursor and store in []messages
	for cursor.Next(ctx) {
		var message domain.Message
		// decode current value in the cursor variable to a message
		cursor.Decode(&message)
		messages = append(messages, message)
	}

	// check if there are any errors with cursor
	if err := cursor.Err(); err != nil {
		restErr := rest.NewInternalServerError("error when searching data of an author", err)
		return nil, restErr
	}

	if len(messages) < 1 {
		return nil, rest.NewInternalServerError("no messages found", err)
	}
	return &messages, nil
}

// Update updates a message
func (repo *mongoRepo) Update(message *domain.Message) (*domain.Response, rest.Err) {

	ctx, cancel := context.WithTimeout(context.Background(), repo.timeout)
	defer cancel()

	collection := repo.client.Database("acm").Collection("messages")

	// filter by
	filter := bson.M{"id": message.ID}
	// update fields
	update := bson.M{
		"$set": bson.M{
			"content":          message.Content,
			"edited_timestamp": message.EditedTimestamp,
		},
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, rest.NewInternalServerError("error updating message in database", err)
	}

	// return response ok
	resp := &domain.Response{Success: "successfully updated message in database"}
	return resp, nil
}

// AddReaction adds a reactions to a message document
func (repo *mongoRepo) AddReaction(r domain.MessageReaction) rest.Err {
	ctx, cancel := context.WithTimeout(context.Background(), repo.timeout)
	defer cancel()

	collection := repo.client.Database("acm").Collection("messages")

	filter := bson.M{"id": r.MessageID}

	mr := domain.MessageReaction{
		UserID: r.UserID,
		Emoji: domain.Emoji{
			Name: r.Emoji.Name,
		},
	}

	update := bson.M{
		"$push": bson.M{"reactions": mr},
	}

	// update db
	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return rest.NewInternalServerError("error when inserting emoji into database", err)
	}
	return nil
}

// Delete deletes a message from the database
func (repo *mongoRepo) Delete(messageID string) rest.Err {
	ctx, cancel := context.WithTimeout(context.Background(), repo.timeout)
	defer cancel()

	collection := repo.client.Database("acm").Collection("messages")

	filter := bson.M{"id": messageID}

	// delete item
	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return rest.NewInternalServerError("error when removing emoji in database", err)
	}
	return nil
}

// RemoveReaction removes a reaction.
func (repo *mongoRepo) DeleteReaction(r domain.MessageReaction) rest.Err {
	ctx, cancel := context.WithTimeout(context.Background(), repo.timeout)
	defer cancel()

	collection := repo.client.Database("acm").Collection("messages")

	filter := bson.M{"id": r.MessageID}
	update := bson.M{
		"$pull": bson.M{
			"reactions": bson.M{
				"user_id":    r.UserID,
				"emoji.name": r.Emoji.Name,
			},
		},
	}

	// update db
	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return rest.NewInternalServerError("error when removing emoji in database", err)
	}
	return nil
}

// InsertLatestPinned inserts a pin into the LatestPin field
// of a channel
func (repo *mongoRepo) InsertLatestPinned(pin *domain.Pin) rest.Err {
	ctx, cancel := context.WithTimeout(context.Background(), repo.timeout)
	defer cancel()

	collection := repo.client.Database("acm").Collection("channels")

	filter := bson.M{"id": pin.ChannelID}
	// if id doesn't exist, create new channel document
	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return rest.NewInternalServerError("error when counting number of docuemnts in database", err)
	}

	// update existing channel document
	if count > 0 {
		update := bson.M{
			"$push": bson.M{
				"latest_pin": bson.M{
					"message":    pin.Message,
					"pinned_at":  pin.PinnedAt,
					"channel_id": pin.ChannelID,
				},
			},
		}

		_, err = collection.UpdateOne(ctx, filter, update)
		if err != nil {
			return rest.NewInternalServerError("error when updating channel in database", err)
		}
	}

	// create new channel document
	channel := &domain.Channel{
		ID:        pin.ChannelID,
		LatestPin: pin,
	}

	// update db
	_, err = collection.InsertOne(ctx, channel)
	if err != nil {
		return rest.NewInternalServerError("error when inserting a new channel into database", err)
	}
	return nil
}

// GetLatestPinned gets the LatestPin from a channel
func (repo *mongoRepo) GetLatestPinned(channelID string) (*domain.Pin, rest.Err) {
	ctx, cancel := context.WithTimeout(context.Background(), repo.timeout)
	defer cancel()

	collection := repo.client.Database("acm").Collection("channels")

	channel := &domain.Channel{}
	result := collection.FindOne(ctx, bson.M{"id": channelID})

	err := result.Decode(channel)
	if err != nil {
		return nil, rest.NewInternalServerError("error when decoding data into object", err)
	}
	return channel.LatestPin, nil
}
