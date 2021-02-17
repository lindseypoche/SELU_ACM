package discord

import "go.mongodb.org/mongo-driver/bson/primitive"

// Channel is a discord channel
type Channel struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name          string             `json:"name,omitempty" bson:"name,omitempty"`
	ChannelID     string             `json:"channel_id,omitempty" bson:"channel_id,omitempty"`
	GuildID       string             `json:"guild_id" bson:"guild_id"`
	LastMessageID string             `json:"last_message_id" bson:"last_message_id"`
}

// Discord Channel
// {
// 	"id": "41771983423143937",
// 	"guild_id": "41771983423143937",
// 	"name": "general",
// 	"type": 0,
// 	"position": 6,
// 	"permission_overwrites": [],
// 	"rate_limit_per_user": 2,
// 	"nsfw": true,
// 	"topic": "24/7 chat about how to gank Mike #2",
// 	"last_message_id": "155117677105512449",
// 	"parent_id": "399942396007890945"
// }
