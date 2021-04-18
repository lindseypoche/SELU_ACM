package listing

import "go.mongodb.org/mongo-driver/bson/primitive"

// A Guild holds all data related to a specific Discord Guild.  Guilds are also
// sometimes referred to as Servers in the Discord client.
// type Guild struct {
// 	ID                     string     `json:"id"`
// 	Name                   string     `json:"name"`
// 	Icon                   string     `json:"icon"`
// 	OwnerID                string     `json:"owner_id"`
// 	Owner                  bool       `json:"owner"`
// 	MemberCount            int        `json:"member_count"`
// 	Roles                  []*Role    `json:"roles"`
// 	Emojis                 []*Emoji   `json:"emojis"`
// 	Members                []*Member  `json:"members"`
// 	Channels               []*Channel `json:"channels"`
// 	Description            string     `json:"description"`
// 	ApproximateMemberCount int        `json:"approximate_member_count"`
// }

// *** Channel has its own document ***
// A Channel holds all data related to an individual Discord channel.
// type Channel struct {
// 	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
// 	DiscordID string             `json:"id" bson:"id"`
// 	GuildID   string             `json:"guild_id,omitempty" bson:"guild_id,omitempty"`
// 	Name      string             `json:"name,omitempty" bson:"name,omitempty"`
// 	Topic     string             `json:"topic,omitempty" bson:"topic,omitempty"`

// 	// If any of these are stored here, only store their reference id.
// 	// They will bloat a database or document.
// 	// Messages  []*Message         `json:"-,omitempty" bson:"-,omitempty"`
// 	// LatestPin *Pin               `json:"latest_pin,omitempty" bson:"latest_pin,omitempty"`
// 	// Pins      []*Pin     `json:"pins,omitempty" bson:"pins,omitempty"`
// }

// *** Message has its own document ***
// A Message stores all data related to a specific Discord message.
type Message struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	DiscordID string             `json:"id" bson:"id"`
	ChannelID string             `json:"channel_id" bson:"channel_id"`

	// ChannelRefID is the mongo reference id referencing a specific channel
	ChannelRefID string `json:"channel_ref_id" bson:"channel_ref_id"`

	GuildID          string             `json:"guild_id,omitempty" bson:"guild_id,omitempty"`
	Content          string             `json:"content" bson:"content"`
	Timestamp        int                `json:"timestamp" bson:"timestamp"`
	EditedTimestamp  int                `json:"edited_timestamp,omitempty" bson:"edited_timestamp,omitempty"`
	MentionRoles     []string           `json:"mention_roles,omitempty" bson:"mention_roles,omitempty"`
	Author           *User              `json:"author,omitempty" bson:"author,omitempty"`
	Attachment       *MessageAttachment `json:"attachments" bson:"attachments,omitempty"`
	Embed            *MessageEmbed      `json:"embeds,omitempty" bson:"embeds,omitemtpy"`
	Mentions         []*User            `json:"mentions,omitempty" bson:"mentions,omitempty"`
	IsPinned         bool               `json:"is_pinned" bson:"is_pinned"`
	MessageReactions MessageReactions   `json:"message_reactions,omitempty" bson:"message_reactions,omitempty"`
	// MessageReference *MessageReference  `json:"message_reference,omitempty" bson:"message_reference,omitempty"`
}

// MessageReference is a reference (ie a reply) to another message.
// A message should not reference another message, else it would be a comment.
// type MessageReference struct {
// 	MessageID string `json:"message_id"`
// 	ChannelID string `json:"channel_id"`
// 	GuildID   string `json:"guild_id,omitempty"`
// }

// Author is a pointer to a User
type Author struct {
	*User
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
	MessageID string `json:"message_id,omitempty" bson:"message_id,omitempty"`
	Emoji     Emoji  `json:"emoji" bson:"emoji"`
	ChannelID string `json:"channel_id,omitempty" bson:"channel_id,omitempty"`
	GuildID   string `json:"guild_id,omitempty" bson:"guild_id,omitempty"`
}

type MessageReactions struct {
	Count     int               `json:"count" bson:"count"`
	Reactions []MessageReaction `json:"reactions,omitempty" bson:"reactions,omitempty"`
}

// Emoji struct holds data related to Emoji's
type Emoji struct {
	ID   string `json:"id,omitempty" bson:"id,omitempty"`
	Name string `json:"name" bson:"name"`
}

// EmpjiUpdate - Unused
// EmojiUpdate is used to update a message's emojis
type EmojiUpdate struct {
	MessageID string `json:"-"`
	Emoji     Emoji  `json:"-"`
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
