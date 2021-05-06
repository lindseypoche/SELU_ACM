package mongo

import (
	"context"
	"log"
	"time"

	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/architecting"
	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/blogging"
	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/commenting"
	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/listing"
	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/reacting"
	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/subscribing"
	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/utils/errors/rest"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

const (
	database = "acm"
	message  = "messages"
	comment  = "comments"
	channel  = "channels"
	member   = "members"
	role     = "roles"
)

var (
	messageCollection = GetClient().Database(database).Collection(message)
	commentCollection = GetClient().Database(database).Collection(comment)
	channelCollection = GetClient().Database(database).Collection(channel)
	memberCollection  = GetClient().Database(database).Collection(member)
	roleCollection    = GetClient().Database(database).Collection(role)
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
func (repo *ListRepo) GetAllMessages() (*[]listing.Message, rest.Err) {

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
	findOptions.SetSort(map[string]int{"timestamp": -1}) // works
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

func (r *ListRepo) GetByStartTime(yesterday int) (*[]listing.Message, rest.Err) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// only get star_time that is less than 24 hours from tomorrow.
	// ie get all future events and events that have expired for 24 hours.
	filter := bson.M{"start_time": bson.M{"$gt": yesterday}}
	findOptions := options.Find()
	findOptions.SetSort(map[string]int{"start_time": 1})

	cursor, err := messageCollection.Find(ctx, filter, findOptions)
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

func (r *ListRepo) GetAllComments(msgID string) (*[]listing.Comment, rest.Err) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	findOptions := options.Find()
	// sort comments by ascending order using timestamp
	findOptions.SetSort(map[string]int{"timestamp": 1}) // works
	// findOptions.SetSort(bson.D{{"timestamp", 1}}) // doesn't work

	// filter based on message_reference id
	filter := bson.M{"message_reference.message_id": msgID}

	cursor, err := commentCollection.Find(ctx, filter, findOptions)
	if err != nil {
		// log error
		return nil, rest.NewInternalServerError("error initializing cursor", err)
	}
	defer cursor.Close(ctx)

	var comments []listing.Comment

	for cursor.Next(ctx) {
		var c listing.Comment
		if err = cursor.Decode(&c); err != nil {
			// log error
			return nil, rest.NewInternalServerError("error decoding a comment", err)
		}

		comments = append(comments, c)
	}

	// check if there are any errors with cursor
	if err = cursor.Err(); err != nil {
		// log error
		return nil, rest.NewInternalServerError("error due to cursor", err)
	}

	if len(comments) < 1 {
		// log error
		return nil, rest.NewNotFoundError("cannot find comments")
	}

	return &comments, nil
}

func (r *ListRepo) GetAllOfficers() (*[]listing.Member, rest.Err) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := memberCollection.Find(ctx, bson.M{})
	if err != nil {
		// log error
		return nil, rest.NewInternalServerError("error initializing cursor", err)
	}
	defer cursor.Close(ctx)

	var members []listing.Member

	for cursor.Next(ctx) {
		var m listing.Member
		if err = cursor.Decode(&m); err != nil {
			// log error
			return nil, rest.NewInternalServerError("error decoding a member", err)
		}

		members = append(members, m)
	}

	// check if there are any errors with cursor
	if err = cursor.Err(); err != nil {
		// log error
		return nil, rest.NewInternalServerError("error due to cursor", err)
	}

	if len(members) < 1 {
		// log error
		return nil, rest.NewNotFoundError("cannot find data")
	}

	return &members, nil
}

func (r *ListRepo) GetActiveOfficers() (*[]listing.Member, rest.Err) {

	return nil, nil
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	// filter := bson.M{"roles": authorUsername}

	// cursor, err := messageCollection.Find(ctx, filter, options.Find().SetSort(map[string]int{"timestamp": 1}))
	// if err != nil {
	// 	return nil, rest.NewInternalServerError("error initializing cursor", err)
	// }
	// defer cursor.Close(ctx)

	// messages := []listing.Message{}
	// // loop through cursor and store in []messages
	// for cursor.Next(ctx) {
	// 	var message listing.Message

	// 	if err = cursor.Decode(&message); err != nil {
	// 		// log error
	// 		return nil, rest.NewInternalServerError("error decoding data into message", err)
	// 	}
	// 	messages = append(messages, message)
	// }

	// // check if there are any errors with cursor
	// if err := cursor.Err(); err != nil {
	// 	// log error
	// 	restErr := rest.NewInternalServerError("error due to cursor", err)
	// 	return nil, restErr
	// }

	// if len(messages) < 1 {
	// 	// log error
	// 	return nil, rest.NewNotFoundError("no messages found")
	// }
	// return &messages, nil
}

func (r *ListRepo) GetUserRoles(id string) (*[]listing.Role, rest.Err) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"user.id": id}

	// get only the roles from the collection
	// findOpts := options.FindOne().SetProjection(bson.M{"roles": 1})

	member := listing.Member{}
	result := memberCollection.FindOne(ctx, filter)

	err := result.Decode(&member)
	log.Println("GetUserRoles: member object\n", member)
	if err != nil {
		return nil, rest.NewInternalServerError("error getting user roles", err)
	}

	if len(*member.Roles) == 0 {
		return nil, rest.NewNotFoundError("error user roles not found")
	}

	return member.Roles, nil
}

func (r *ListRepo) GetAllRoles() (*[]listing.Role, rest.Err) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := roleCollection.Find(ctx, bson.M{})
	if err != nil {
		// log error
		return nil, rest.NewInternalServerError("error initializing cursor", err)
	}
	defer cursor.Close(ctx)

	roles := []listing.Role{}

	for cursor.Next(ctx) {
		var role listing.Role
		if err = cursor.Decode(&role); err != nil {
			// log error
			return nil, rest.NewInternalServerError("error decoding a role", err)
		}
		roles = append(roles, role)
	}

	result, err := roleCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, rest.NewInternalServerError("error getting user roles", err)
	}

	if err = result.Err(); err != nil {
		return nil, rest.NewInternalServerError("error from cursor", err)
	}

	// check if there are any errors with cursor
	// if err = cursor.Err(); err != nil {
	// 	// log error
	// 	return nil, rest.NewInternalServerError("error due to cursor", err)
	// }

	if len(roles) == 0 {
		return nil, rest.NewNotFoundError("error user roles not found")
	}

	return &roles, nil
}

func (r *ListRepo) GetChannel(id string) (*listing.Channel, rest.Err) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	channel := &listing.Channel{}
	result := channelCollection.FindOne(ctx, bson.M{"id": id})

	err := result.Decode(channel)
	if err != nil {
		return nil, rest.NewInternalServerError("error decoding channel", err)
	}
	return channel, nil
}

// BlogRepo
type BlogRepo struct{}

// SaveMessage ( current status: ✅ )
// Attempts to save a message into the database
func (r *BlogRepo) SaveMessage(message *blogging.Message) rest.Err {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// var cid, mid bson.ObjectId

	// // retreive only the _id
	// findOpts := options.FindOne().SetProjection(bson.M{"_id": 1})

	// // get Channel Ref ID and Member Ref ID
	// err := channelCollection.FindOne(ctx, bson.M{"id": message.ChannelID}, findOpts).Decode(&cid)
	// if err != nil {
	// 	if err == mongo.ErrNoDocuments {
	// 		log.Println(rest.NewNotFoundError("channel not found")) // warning only
	// 	}
	// 	log.Println(rest.NewInternalServerError("error searching for channel", err)) // warning only
	// }

	// err = memberCollection.FindOne(ctx, bson.M{"user.id": message.Author.ID}, findOpts).Decode(&mid)
	// if err != nil {
	// 	if err == mongo.ErrNoDocuments {
	// 		log.Println(rest.NewNotFoundError("message not member")) // warning only
	// 	}
	// 	log.Println(rest.NewInternalServerError("error searching for member", err)) // warning only
	// }

	// // store both references in new message
	// message.ChannelRefID = cid
	// message.MemberRefID = mid

	_, err = messageCollection.InsertOne(ctx, *message)
	if err != nil {
		// log error
		return rest.NewInternalServerError("error inserting message", err)
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

	// delete from messages collection
	result, err := messageCollection.DeleteOne(ctx, filter)
	if err != nil {
		// log error
		return rest.NewInternalServerError("error when deleting message", err)
	}

	if result.DeletedCount == 0 {
		return rest.NewNotFoundError("error comment not found to delete")
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

// GetReactions (postponed)
func (r *ReactRepo) GetReactions(string) (*reacting.MessageReactions, rest.Err) {
	return nil, rest.NewStatusNotImplemented("get reactions not implemented")
}

type CommentRepo struct{}

func (r *CommentRepo) SaveComment(comment *commenting.Comment) rest.Err {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := commentCollection.InsertOne(ctx, *comment)
	if err != nil {
		// log error
		return rest.NewInternalServerError("error inserting comment", err)
	}

	return nil
}

func (r *CommentRepo) EditComment(comment *commenting.Comment) rest.Err {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// filter by the comment's discord id
	filter := bson.M{"id": comment.DiscordID}
	// update all fields except the original timestamp
	update := bson.M{
		"$set": bson.M{
			"content":          comment.Content,
			"edited_timestamp": comment.EditedTimestamp,
			// "message_reference": comment.MessageReference,
		},
	}

	result, err := commentCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		// log error
		return rest.NewInternalServerError("error when attempting to update comment", err)
	}

	if result.ModifiedCount == 0 {
		return rest.NewNotFoundError("could not find any matching documents in comments collection to edit")
	}

	return nil
}

func (r *CommentRepo) DeleteComment(id string) rest.Err {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"id": id}

	// delete from comments collection
	result, err := commentCollection.DeleteOne(ctx, filter)
	if err != nil {
		// log error
		return rest.NewInternalServerError("error when deleting comment", err)
	}

	if result.DeletedCount == 0 {
		return rest.NewNotFoundError("could not find any matching documents in comments collection to delete")
	}

	return nil
}

// todo: implement
// func (r *CommentRepo) DeleteAllComments(refID string) rest.Err {

// }

// Subscribing
type SubscribeRepo struct{}

func (r *SubscribeRepo) SaveMember(member *subscribing.Member) rest.Err {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := memberCollection.InsertOne(ctx, *member)
	if err != nil {
		// log error
		return rest.NewInternalServerError("error saving member", err)
	}
	return nil
}

func (r *SubscribeRepo) UpdateMember(member *subscribing.Member) rest.Err {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"user.id": member.User.ID}

	update := bson.M{
		"$set": bson.M{
			"nick":    member.Nick,
			"user":    member.User,
			"roles":   member.Roles,
			"content": member.Content,
		},
	}

	result, err := memberCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		// log error
		return rest.NewInternalServerError("error updating member", err)
	}

	if result.ModifiedCount == 0 {
		return rest.NewNotFoundError("error member not found")
	}

	return nil
}

func (r *SubscribeRepo) UpdateOfficerContent(member *subscribing.Member, id string) rest.Err {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"user.id": id}

	update := bson.M{
		"$set": bson.M{
			"content": member.Content,
		},
	}

	result, err := memberCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		// log error
		return rest.NewInternalServerError("error updating officer content", err)
	}

	if result.ModifiedCount == 0 {
		return rest.NewNotFoundError("error officer not found")
	}

	return nil
}

func (r *SubscribeRepo) DeleteOfficerContent(id string) rest.Err {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"user.id": id}

	update := bson.M{
		"$unset": bson.M{
			"content": "",
		},
	}

	result, err := memberCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		// log error
		return rest.NewInternalServerError("error deleting officer content", err)
	}

	if result.ModifiedCount == 0 {
		return rest.NewNotFoundError("error officer not found")
	}

	return nil
}

// NOTICE: if you do not want a member to be deleted, like an officer,
// then you should assign them an acm_alumi role before removing other
// acm roles.
func (r *SubscribeRepo) DeleteMember(id string) rest.Err {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"user.id": id}
	result, err := memberCollection.DeleteOne(ctx, filter)
	if err != nil {
		// log error
		return rest.NewInternalServerError("error deleting member", err)
	}

	if result.DeletedCount == 0 {
		return rest.NewNotFoundError("error member id not found")
	}

	return nil
}

// ArchitectRepo
type ArchitectRepo struct{}

// DO NOT USE.
func (r *ArchitectRepo) SaveRole(role *architecting.Role) rest.Err {
	return nil
}

func (r *ArchitectRepo) SaveChannel(channel *architecting.Channel) rest.Err {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := channelCollection.InsertOne(ctx, *channel)
	if err != nil {
		// log error
		return rest.NewInternalServerError("error inserting channel", err)
	}
	return nil
}

func (r *ArchitectRepo) UpdateRole(role *architecting.Role) rest.Err {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"id": role.DiscordID}

	update := bson.M{
		"$set": bson.M{
			"id":   role.DiscordID,
			"name": role.Name,
		},
	}

	result, err := roleCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		// log error
		return rest.NewInternalServerError("error updating role", err)
	}

	// if could not modify role, it doesn't exist. so attempt to save the role.
	if result.MatchedCount == 0 {
		_, err = roleCollection.InsertOne(ctx, *role)
		if err != nil {
			// log error
			return rest.NewInternalServerError("error inserting role", err)
		} else if err == nil {
			// successfully saved role
			return nil
		}

		return rest.NewNotFoundError("error role id not found to update")
	}
	return nil
}

func (r *ArchitectRepo) DeleteRole(id string) rest.Err {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"id": id}
	result, err := roleCollection.DeleteOne(ctx, filter)
	if err != nil {
		// log error
		return rest.NewInternalServerError("error deleting role", err)
	}
	if result.DeletedCount == 0 {
		return rest.NewNotFoundError("error role id not found to delete")
	}
	return nil
}
