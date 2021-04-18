package commenting

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
	MessageID string `json:"message_id" bson:"message_id,omitempty"`
	ChannelID string `json:"channel_id" bson:"channel_id,omitempty"`
}

// Author is a pointer to a User
type Author struct {
	*User
}

// A User stores all data for an individual Discord user.
type User struct {
	ID            string `json:"id" bson:"id"`
	Username      string `json:"username" bson:"username"`
	Discriminator string `json:"discriminator" bson:"discriminator"`
	Avatar        Avatar `json:"avatar,omitempty" bson:"avatar,omitempty"`
	Email         string `json:"email,omitempty" bson:"email,omitempty"`
}

// *** Member has its own document ***
// A Member stores user information for Guild members. A guild
// member represents a certain user's presence in a guild.
type Member struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	GuildID string             `json:"guild_id,omitempty" bson:"guild_id,omitempty"`

	// JoinedAt is a unix timestamp
	JoinedAt int    `json:"joined_at" bson:"joined_at"`
	Nick     string `json:"nick" bson:"nick"`
	User     *User  `json:"user" bson:"user"`
	Roles    []Role `json:"roles,omitempty" bson:"roles,omitempty"`

	// OfficerStatus includes: active, inactive, nil
	OfficerStatus string `json:"officer_status,omitempty" bson:"officer_status,omitempty"`
}

// Avatar data of a user's discord avatar
type Avatar struct {
	ID       string `json:"id" bson:"id"`
	ImageURL string `json:"image_url" bson:"image_url"`
}

// A Role stores information about Discord guild member roles.
type Role struct {
	ID   string `json:"id,omitempty" bson:"id,omitempty"`
	Name string `json:"name" bson:"name"`
}
