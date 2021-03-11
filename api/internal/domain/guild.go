package domain

/*
Discord objects and fields
*/

// A Guild holds all data related to a specific Discord Guild.  Guilds are also
// sometimes referred to as Servers in the Discord client.
type Guild struct {
	ID                     string     `json:"id"`
	Name                   string     `json:"name"`
	Icon                   string     `json:"icon"`
	OwnerID                string     `json:"owner_id"`
	Owner                  bool       `json:"owner"`
	MemberCount            int        `json:"member_count"`
	Roles                  []*Role    `json:"roles"`
	Emojis                 []*Emoji   `json:"emojis"`
	Members                []*Member  `json:"members"`
	Channels               []*Channel `json:"channels"`
	Description            string     `json:"description"`
	ApproximateMemberCount int        `json:"approximate_member_count"`
}

// A Channel holds all data related to an individual Discord channel.
type Channel struct {
	ID        string     `json:"id"`
	GuildID   string     `json:"guild_id,omitempty" bson:"guild_id,omitempty"`
	Name      string     `json:"name,omitempty" bson:"name,omitempty"`
	Topic     string     `json:"topic,omitempty" bson:"topic,omitempty"`
	Messages  []*Message `json:"-,omitempty" bson:"-,omitempty"`
	LatestPin *Pin       `json:"latest_pin,omitempty" bson:"latest_pin,omitempty"`
	// Pins      []*Pin     `json:"pins,omitempty" bson:"pins,omitempty"`
}

// A Message stores all data related to a specific Discord message.
type Message struct {
	ID              string             `json:"id" bson:"id"`
	ChannelID       string             `json:"channel_id" bson:"channel_id"`
	GuildID         string             `json:"guild_id,omitempty" bson:"guild_id,omitempty"`
	Content         string             `json:"content" bson:"content"`
	Timestamp       int                `json:"timestamp" bson:"timestamp"`
	EditedTimestamp int                `json:"edited_timestamp,omitempty" bson:"edited_timestamp,omitempty"`
	MentionRoles    []string           `json:"mention_roles,omitempty" bson:"mention_roles,omitempty"`
	Author          *User              `json:"author"`
	Attachment      *MessageAttachment `json:"attachments" bson:"attachments,omitempty"`
	Embeds          *MessageEmbed      `json:"embeds,omitempty" bson:"embeds,omitemtpy"`
	Mentions        []*User            `json:"mentions,omitempty" bson:"mentions,omitempty"`
	Pinned          bool               `json:"pinned,omitempty" bson:"pinned,omitempty"`
	Reactions       []MessageReaction  `json:"reactions,omitempty" bson:"reactions,omitempty"`
}

// Author is a pointer to a User
type Author struct {
	*User
}

// A User stores all data for an individual Discord user.
type User struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
	Avatar        Avatar `json:"avatar"`
	Email         string `json:"email"`
}

// A Member stores user information for Guild members. A guild
// member represents a certain user's presence in a guild.
type Member struct {
	GuildID  string    `json:"guild_id"`
	JoinedAt Timestamp `json:"joined_at"`
	Nick     string    `json:"nick"`
	User     *User     `json:"user"`
	Roles    []string  `json:"roles"`
}

// Avatar ...
type Avatar struct {
	ID       string `json:"id"`
	ImageURL string `json:"image_url"`
}

// A MessageAttachment stores data for message attachments.
type MessageAttachment struct {
	ID       string `json:"id"`
	URL      string `json:"url"`
	Filename string `json:"filename"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	Size     int    `json:"size"`
}

// MessageReaction ...
type MessageReaction struct {
	UserID    string `json:"user_id" bson:"user_id"`
	MessageID string `json:"message_id" bson:"message_id,omitempty"`
	Emoji     Emoji  `json:"emoji" bson:"emoji"`
	ChannelID string `json:"channel_id" bson:"channel_id,omitempty"`
	GuildID   string `json:"guild_id,omitempty" bson:"guild_id,omitempty"`
}

// Emoji struct holds data related to Emoji's
type Emoji struct {
	ID    string `json:"id,omitempty" bson:"id,omitempty"`
	Name  string `json:"name" bson:"name"`
	Count int    `json:"count" bson:"count,omitempty"`
}

// EmojiUpdate is used to update a message's emojis
type EmojiUpdate struct {
	MessageID string `json:"-"`
	Emoji     Emoji  `json:"-"`
}

// An MessageEmbed stores data for message embeds.
type MessageEmbed struct {
	URL         string              `json:"url,omitempty"`
	Type        string              `json:"type,omitempty"`
	Title       string              `json:"title,omitempty"`
	Description string              `json:"description,omitempty"`
	Timestamp   int                 `json:"timestamp,omitempty"`
	Color       int                 `json:"color,omitempty"`
	Image       *MessageEmbedImage  `json:"image,omitempty"`
	Video       *MessageEmbedVideo  `json:"video,omitempty"`
	Author      *MessageEmbedAuthor `json:"author,omitempty"`
}

// MessageEmbedImage is a part of a MessageEmbed struct.
type MessageEmbedImage struct {
	URL      string `json:"url,omitempty"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Width    int    `json:"width,omitempty"`
	Height   int    `json:"height,omitempty"`
}

// MessageEmbedVideo is a part of a MessageEmbed struct.
type MessageEmbedVideo struct {
	URL    string `json:"url,omitempty"`
	Width  int    `json:"width,omitempty"`
	Height int    `json:"height,omitempty"`
}

// MessageEmbedAuthor is a part of a MessageEmbed struct.
type MessageEmbedAuthor struct {
	URL          string `json:"url,omitempty"`
	Name         string `json:"name,omitempty"`
	IconURL      string `json:"icon_url,omitempty"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

// A Role stores information about Discord guild member roles.
type Role struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Pin stores data about pinned messages in a channel
type Pin struct {
	Message   *Message `json:"message" bson:"message"`
	PinnedAt  int      `json:"pinned_at" bson:"pinned_at"`
	ChannelID string   `json:"channel_id,omitempty" bson:"channel_id,omitempty"`
}

// Timestamp stores a timestamp, as sent by the Discord API.
type Timestamp int
