package architecting

type Role struct {
	DiscordID string `json:"id" bson:"id"`
	Name      string `json:"name" bson:"name"`
}
