package listing

import "go.mongodb.org/mongo-driver/bson/primitive"

type Role struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	DiscordID string             `json:"id" bson:"id"`
	Name      string             `json:"name" bson:"name"`
}
