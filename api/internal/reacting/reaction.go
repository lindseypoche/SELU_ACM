package reacting

// MessageReaction ...
type MessageReaction struct {
	UserID    string `json:"user_id" bson:"user_id"`
	MessageID string `json:"message_id" bson:"message_id"`
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
	ID    string `json:"id,omitempty" bson:"id,omitempty"`
	Name  string `json:"name" bson:"name"`
	Count int    `json:"count,omitempty" bson:"count,omitempty"`
}
