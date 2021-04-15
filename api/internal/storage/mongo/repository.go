package mongo

import (
	"context"
	"log"
	"time"

	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/blogging"
	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/listing"
	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/reacting"
	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/utils/errors/rest"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

const (
	database = "acm"
	message  = "messages"
	channel  = "channels"
	member   = "members"
)

var (
	messageCollection = GetClient().Database(database).Collection(message)
	channelCollection = GetClient().Database(database).Collection(channel)
	memberCollection  = GetClient().Database(database).Collection(member)
)

// ListRepo
// *could* embed mongo client into struct
type ListRepo struct{}

// GetByID ( current status: ✅ )
// GetByID attempts to get a messages by id from the database
func (repo *ListRepo) GetByID(messageID string) (*listing.Message, rest.Err) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	message := &listing.Message{}
	result := messageCollection.FindOne(ctx, bson.M{"id": messageID})

	err := result.Decode(message)
	if err != nil {
		return nil, rest.NewInternalServerError("error decoding message", err)
	}
	return message, nil
}

// GetAll ( current status: ✅ )
// GetAll gets all messages in the database
func (repo *ListRepo) GetAll() (*[]listing.Message, rest.Err) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := messageCollection.Find(ctx, bson.M{})
	if err != nil {
		// log error
		return nil, rest.NewInternalServerError("error initializing cursor", err)
	}
	defer cursor.Close(ctx)

	var messages []listing.Message

	for cursor.Next(ctx) {
		var m listing.Message
		if err = cursor.Decode(&m); err != nil {
			// log error
			return nil, rest.NewInternalServerError("error decoding a message", err)
		}

		messages = append(messages, m)
	}

	// check if there are any errors with cursor
	if err = cursor.Err(); err != nil {
		// log error
		return nil, rest.NewInternalServerError("error due to cursor", err)
	}

	if len(messages) < 1 {
		// log error
		return nil, rest.NewNotFoundError("cannot find data")
	}

	return &messages, nil
}

// GetByUsername ( current status: ✅ )
// GetByAuthor gets all the posts created by an author
func (repo *ListRepo) GetByUsername(authorUsername string) (*[]listing.Message, rest.Err) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"author.username": authorUsername}

	cursor, err := messageCollection.Find(ctx, filter, options.Find().SetSort(map[string]int{"timestamp": 1}))
	if err != nil {
		return nil, rest.NewInternalServerError("error initializing cursor", err)
	}
	defer cursor.Close(ctx)

	messages := []listing.Message{}
	// loop through cursor and store in []messages
	for cursor.Next(ctx) {
		var message listing.Message

		if err = cursor.Decode(&message); err != nil {
			// log error
			return nil, rest.NewInternalServerError("error decoding data into message", err)
		}
		messages = append(messages, message)
	}

	// check if there are any errors with cursor
	if err := cursor.Err(); err != nil {
		// log error
		restErr := rest.NewInternalServerError("error due to cursor", err)
		return nil, restErr
	}

	if len(messages) < 1 {
		// log error
		return nil, rest.NewNotFoundError("no messages found")
	}
	return &messages, nil
}

// GetPinnedMessageByID ( current status: ✅ )
// GetPinnedMessage searching for a message by its messageID and returns it
// if it exists.
func (repo *ListRepo) GetPinnedMessageByID(messageID string) (*listing.Message, rest.Err) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	message := &listing.Message{}
	filter := bson.M{"id": messageID}

	result := messageCollection.FindOne(ctx, filter)
	err := result.Decode(message)
	if err != nil {
		return nil, rest.NewInternalServerError("error decoding message", err)
	}

	if result == nil || !message.IsPinned {
		return nil, rest.NewNotFoundError("message not found")
	}

	return message, nil
}

// GetLatestPinned ( current status: ✅ )
// GetLatestPinned gets the most recent message that was pinned in a channel
func (repo *ListRepo) GetLatestPinned(channelID string) (*listing.Message, rest.Err) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{
		"channel_id": channelID,
		"is_pinned":  true,
	}

	findOptions := options.Find()
	findOptions.SetSort(map[string]int{"_id": -1})
	findOptions.SetLimit(1)

	cursor, err := messageCollection.Find(ctx, filter, findOptions)

	m := &listing.Message{}
	for cursor.Next(ctx) {
		cursor.Decode(m)
	}

	if err = cursor.Err(); err != nil {
		// log error
		return nil, rest.NewInternalServerError("error due to cursor", err)
	}
	if m == nil {
		return nil, rest.NewNotFoundError("message not found")
	}

	return m, nil
}

func (repo *ListRepo) GetAllPinnedMessages() (*[]listing.Message, rest.Err) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"is_pinned": true}

	cursor, err := messageCollection.Find(ctx, filter)
	if err != nil {
		// log error
		return nil, rest.NewInternalServerError("error initializing cursor", err)
	}
	defer cursor.Close(ctx)

	var messages []listing.Message

	for cursor.Next(ctx) {
		var m listing.Message
		if err = cursor.Decode(&m); err != nil {
			// log error
			return nil, rest.NewInternalServerError("error decoding a message", err)
		}

		messages = append(messages, m)
	}

	// check if there are any errors with cursor
	if err = cursor.Err(); err != nil {
		// log error
		return nil, rest.NewInternalServerError("error due to cursor", err)
	}

	if len(messages) < 1 {
		// log error
		return nil, rest.NewNotFoundError("cannot find data")
	}

	return &messages, nil
}

// BlogRepo
type BlogRepo struct{}

// SaveMessage ( current status: ✅ )
// Attempts to save a message into the database
func (r *BlogRepo) SaveMessage(message *blogging.Message) rest.Err {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := messageCollection.InsertOne(ctx, *message)
	if err != nil {
		// log error
		return rest.NewInternalServerError("error inserting data", err)
	}

	return nil
}

// EditMessage ( current status: ✅ )
func (r *BlogRepo) EditMessage(message *blogging.Message) (*listing.Message, rest.Err) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// filter by the message's discord id
	filter := bson.M{"id": message.DiscordID}
	// update all fields except the original timestamp
	update := bson.M{
		"$set": bson.M{
			"content":          message.Content,
			"edited_timestamp": message.EditedTimestamp,
		},
	}

	// options.ReturnDocument = 1 will return the item after it is stored
	result := messageCollection.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(1))
	if result == nil {
		// log error
		log.Println("message not found")
		return nil, rest.NewNotFoundError("message not found")
	}

	m := &listing.Message{}
	err := result.Decode(m)
	if err != nil {
		log.Println("error decoding message")
		return nil, rest.NewInternalServerError("error decoding message", err)
	}

	return m, nil
}

// DeleteMessage ( current status: ✅ )
func (r *BlogRepo) DeleteMessage(id string) rest.Err {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"id": id}

	// delete from messages document db
	_, err := messageCollection.DeleteOne(ctx, filter)
	if err != nil {
		// log error
		return rest.NewInternalServerError("error when deleting message", err)
	}

	return nil
}

// UpdatePinMessage ( current status: ✅ )
// UpdatePinMessage updates a message's is_pinned property
// and returns the new updated message
func (r *BlogRepo) UpdatePinMessage(id string, isPinned bool) (*listing.Message, rest.Err) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// filter by the message's discord id
	filter := bson.M{"id": id}

	update := bson.M{"$set": bson.M{"is_pinned": isPinned}}

	result := messageCollection.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(1))
	if result == nil {
		// log error
		log.Println("message not found")
		return nil, rest.NewNotFoundError("message not found")
	}

	m := &listing.Message{}
	err := result.Decode(m)
	if err != nil {
		log.Println("error decoding message")
		return nil, rest.NewInternalServerError("error decoding message", err)
	}

	return m, nil
}

// *** Reacting repo ***
type ReactRepo struct{}

// SaveReaction ( current status: ✅ )
func (r *ReactRepo) SaveReaction(mr reacting.MessageReaction) rest.Err {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"id": mr.MessageID}

	update := bson.M{
		"$push": bson.M{"message_reactions.reactions": mr},
		"$inc":  bson.M{"message_reactions.count": 1},
	}

	// update db
	_, err := messageCollection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	if err != nil {
		// log error
		log.Println("DB: error when inserting emoji")
		return rest.NewInternalServerError("error when inserting emoji", err)
	}
	return nil
}

// GetReactions ( current status: ❌ )
// GetReactions is not necessary yet as reactions are sent with a message.
func (r *ReactRepo) GetReactions(string) (*reacting.MessageReactions, rest.Err) {
	return nil, nil
}

// DeleteReaction ( current status: ✅ )
func (r *ReactRepo) DeleteReaction(mr reacting.MessageReaction) rest.Err {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"id": mr.MessageID}

	update := bson.M{
		"$pull": bson.M{
			"message_reactions.reactions": bson.M{
				"user_id":    mr.UserID,
				"emoji.name": mr.Emoji.Name,
			},
		},
		"$inc": bson.M{"message_reactions.count": -1},
	}

	// update db
	_, err := messageCollection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	if err != nil {
		// log error
		return rest.NewInternalServerError("error when removing emoji", err)
	}

	return nil
}
