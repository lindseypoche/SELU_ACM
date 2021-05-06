package listing

import "go.mongodb.org/mongo-driver/bson/primitive"

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
	ID      primitive.ObjectID `json:"_id" bson:"_id"`
	GuildID string             `json:"guild_id,omitempty" bson:"guild_id,omitempty"`

	// JoinedAt is a unix timestamp
	JoinedAt int    `json:"joined_at" bson:"joined_at"`
	Nick     string `json:"nick" bson:"nick"`
	User     *User  `json:"user" bson:"user"`

	// MemberRole is the roles of a member
	Roles *[]Role `json:"roles,omitempty" bson:"roles,omitempty"`

	IsOfficer bool `json:"is_officer,omitempty" bson:"is_officer,omitempty"`
}

// Avatar data of a user's discord avatar
type Avatar struct {
	ID       string `json:"id" bson:"id"`
	ImageURL string `json:"image_url" bson:"image_url"`
}
