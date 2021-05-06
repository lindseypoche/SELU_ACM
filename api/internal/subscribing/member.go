package subscribing

import "go.mongodb.org/mongo-driver/bson/primitive"

type IsActive bool

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

	// Roles uses the roles primitive.ObjectID in the roles collection
	// to check if the user is currently active in that role.
	// If IsActive == false, then user was affiliated with the role in the past,
	// but currently is not.
	// Roles map[primitive.ObjectID]IsActive `json:"roles,omitempty" bson:"roles,omitempty"`
	Roles *[]Role `json:"roles,omitempty" bson:"roles,omitempty"`

	// Officer content
	Content *Content `json:"content,omitempty" bson:"content,omitempty"`
}

type Content struct {
	Chair string `json:"position" bson:"position"`
	Text  string `json:"text" bson:"text"`
	Photo string `json:"photo" bson:"photo"`
}

// Avatar data of a user's discord avatar
type Avatar struct {
	ID       string `json:"id" bson:"id"`
	ImageURL string `json:"image_url" bson:"image_url"`
}
