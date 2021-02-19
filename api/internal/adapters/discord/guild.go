package discord

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
	ID            string     `json:"id"`
	GuildID       string     `json:"guild_id"`
	Name          string     `json:"name"`
	Topic         string     `json:"topic"`
	Messages      []*Message `json:"-"`
	OwnerID       string     `json:"owner_id"`
	ApplicationID string     `json:"application_id"`
}

// A Message stores all data related to a specific Discord message.
type Message struct {
	ID              string        `json:"id"`
	ChannelID       string        `json:"channel_id"`
	GuildID         string        `json:"guild_id,omitempty"`
	Content         string        `json:"content"`
	Timestamp       Timestamp     `json:"timestamp"`
	EditedTimestamp Timestamp     `json:"edited_timestamp"`
	MentionRoles    []string      `json:"mention_roles"`
	Author          *User         `json:"author"`
	Embeds          *MessageEmbed `json:"embeds"`
	Mentions        []*User       `json:"mentions"`
	Member          *Member       `json:"member"`
}

// A User stores all data for an individual Discord user.
type User struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
	Avatar        string `json:"avatar"`
	Email         string `json:"email"`
}

// A Member stores user information for Guild members. A guild
// member represents a certain user's presence in a guild.
type Member struct {
	GuildID  string    `json:"guild_id"`
	JoinedAt Timestamp `json:"joined_at"`
	Nick     string    `json:"nick"`
	Deaf     bool      `json:"deaf"`
	User     *User     `json:"user"`
	Roles    []string  `json:"roles"`
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

// MessageReactions holds a reactions object for a message.
type MessageReactions struct {
	Count int    `json:"count"`
	Emoji *Emoji `json:"emoji"`
}

// MessageReaction ...
type MessageReaction struct {
	UserID    string `json:"user_id"`
	MessageID string `json:"message_id"`
	Emoji     Emoji  `json:"emoji"`
	ChannelID string `json:"channel_id"`
	GuildID   string `json:"guild_id,omitempty"`
}

// Emoji struct holds data related to Emoji's
type Emoji struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// An MessageEmbed stores data for message embeds.
type MessageEmbed struct {
	URL         string              `json:"url,omitempty"`
	Type        string              `json:"type,omitempty"`
	Title       string              `json:"title,omitempty"`
	Description string              `json:"description,omitempty"`
	Timestamp   Timestamp           `json:"timestamp,omitempty"`
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

// Timestamp stores a timestamp, as sent by the Discord API.
type Timestamp string
