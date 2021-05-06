package subscribing

import "go.mongodb.org/mongo-driver/bson/primitive"

// A Role stores information about Discord guild member roles.
type Role struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	DiscordID string             `json:"id" bson:"id"`
	Name      string             `json:"name" bson:"name"`
}
