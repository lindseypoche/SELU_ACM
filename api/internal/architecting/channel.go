package architecting

// A Channel holds all data related to an individual Discord channel.
type Channel struct {
	DiscordID  string `json:"id" bson:"id"`
	GuildID    string `json:"guild_id,omitempty" bson:"guild_id,omitempty"`
	Name       string `json:"name,omitempty" bson:"name,omitempty"`
	Topic      string `json:"topic,omitempty" bson:"topic,omitempty"`
	Collection string `json:"collection" bson:"collection"`
}
