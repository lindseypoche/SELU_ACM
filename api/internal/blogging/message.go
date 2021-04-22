package blogging

import "go.mongodb.org/mongo-driver/bson/primitive"

// *** Message has its own document ***
// A Message stores all data related to a specific Discord message.
type Message struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	DiscordID string             `json:"id" bson:"id"`
	ChannelID string             `json:"channel_id" bson:"channel_id"`

	// ChannelRefID is the mongo reference id referencing a specific channel
	ChannelRefID string `json:"channel_ref_id" bson:"channel_ref_id"`
	GuildID      string `json:"guild_id,omitempty" bson:"guild_id,omitempty"`

	// Event fields
	StartTime int    `json:"start_time" bson:"start_time"`
	Title     string `json:"title" bson:"title"`
	Content   string `json:"content" bson:"content"`

	Timestamp        int                `json:"timestamp" bson:"timestamp"`
	EditedTimestamp  int                `json:"edited_timestamp,omitempty" bson:"edited_timestamp,omitempty"`
	MentionRoles     []string           `json:"mention_roles,omitempty" bson:"mention_roles,omitempty"`
	Author           *User              `json:"author,omitempty" bson:"author,omitempty"`
	Attachment       *MessageAttachment `json:"attachments" bson:"attachments,omitempty"`
	Embed            *MessageEmbed      `json:"embeds,omitempty" bson:"embeds,omitemtpy"`
	Mentions         []*User            `json:"mentions,omitempty" bson:"mentions,omitempty"`
	IsPinned         bool               `json:"is_pinned" bson:"is_pinned"`
	MessageReactions MessageReactions   `json:"message_reactions,omitempty" bson:"message_reactions,omitempty"`
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

// A MessageAttachment stores data for message attachments.
type MessageAttachment struct {
	ID       string `json:"id" bson:"id"`
	URL      string `json:"url" bson:"url"`
	Filename string `json:"filename" bson:"filename"`
	Width    int    `json:"width" bson:"width"`
	Height   int    `json:"height" bson:"height"`
	Size     int    `json:"size" bson:"size"`
}

// MessageReaction ...
type MessageReaction struct {
	UserID    string `json:"user_id" bson:"user_id"`
	MessageID string `json:"message_id" bson:"message_id,omitempty"`
	Emoji     Emoji  `json:"emoji" bson:"emoji"`
	ChannelID string `json:"channel_id" bson:"channel_id,omitempty"`
	GuildID   string `json:"guild_id,omitempty" bson:"guild_id,omitempty"`
}

type MessageReactions struct {
	Count     int               `json:"count" bson:"count"`
	Reactions []MessageReaction `json:"reactions,omitempty" bson:"reactions,omitempty"`
}

// Emoji struct holds data related to Emoji's
type Emoji struct {
	ID    string `json:"id,omitempty" bson:"id,omitempty"`
	Name  string `json:"name" bson:"name"`
	Count int    `json:"count" bson:"count,omitempty"`
}

// An MessageEmbed stores data for message embeds.
type MessageEmbed struct {
	URL         string              `json:"url" bson:"url"`
	Type        string              `json:"type,omitempty" bson:"type,omitempty"`
	Title       string              `json:"title,omitempty" bson:"title,omitempty"`
	Description string              `json:"description,omitempty" bson:"description,omitempty"`
	Timestamp   int                 `json:"timestamp" bson:"timestamp"`
	Color       int                 `json:"color,omitempty" bson:"color,omitempty"`
	Image       *MessageEmbedImage  `json:"image,omitempty" bson:"image,omitempty"`
	Video       *MessageEmbedVideo  `json:"video,omitempty" bson:"video,omitempty"`
	Author      *MessageEmbedAuthor `json:"author,omitempty" bson:"author,omitempty"`
}

// MessageEmbedImage is a part of a MessageEmbed struct.
type MessageEmbedImage struct {
	URL      string `json:"url" bson:"url"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Width    int    `json:"width,omitempty" bson:"width,omitempty"`
	Height   int    `json:"height,omitempty" bson:"height,omitempty"`
}

// MessageEmbedVideo is a part of a MessageEmbed struct.
type MessageEmbedVideo struct {
	URL    string `json:"url" bson:"url"`
	Width  int    `json:"width,omitempty" bson:"width"`
	Height int    `json:"height,omitempty" bson:"height"`
}

// MessageEmbedAuthor is a part of a MessageEmbed struct.
type MessageEmbedAuthor struct {
	URL          string `json:"url" bson:"url"`
	Name         string `json:"name" bson:"name"`
	IconURL      string `json:"icon_url" bson:"icon_url"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty" bson:"proxy_icon_url,omitempty"`
}

// A Role stores information about Discord guild member roles.
type Role struct {
	ID   string `json:"id,omitempty" bson:"id,omitempty"`
	Name string `json:"name" bson:"name"`
}

// Pin stores data about pinned messages in a channel
type Pin struct {
	Message  *Message `json:"message" bson:"message"`
	PinnedAt int      `json:"pinned_at" bson:"pinned_at"`
}
