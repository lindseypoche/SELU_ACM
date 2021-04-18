package listing

import "go.mongodb.org/mongo-driver/bson/primitive"

type Comment struct {
	ID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	DiscordID       string             `json:"id" bson:"id"`
	ChannelID       string             `json:"channel_id" bson:"channel_id"`
	Content         string             `json:"content" bson:"content"`
	Timestamp       int                `json:"timestamp" bson:"timestamp"`
	EditedTimestamp int                `json:"edited_timestamp,omitempty" bson:"edited_timestamp,omitempty"`
	Author          *User              `json:"author,omitempty" bson:"author,omitempty"`
	// Mentions        []*User            `json:"mentions,omitempty" bson:"mentions,omitempty"`

	// MessageReference is a reference to a message like an event.
	MessageReference *MessageReference `json:"message_reference,omitempty" bson:"message_reference,omitempty"`
}

// MessageReference is a reference (ie a reply) to another message.
type MessageReference struct {
	MessageID string `json:"message_id,omitempty" bson:"message_id,omitempty"`
	ChannelID string `json:"channel_id,omitempty" bson:"channel_id,omitempty"`
}
