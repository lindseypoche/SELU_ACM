package blogs

import "go.mongodb.org/mongo-driver/bson/primitive"

// Blog ...
type Blog struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title,omitempty" bson:"title,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Image       string             `json:"image" bson:"image"`
	Author      Author             `json:"author" bson:"author"`
	DateCreated string             `json:"date_created"`
	DateUpdated string             `json:"date_updated"`
}

// Author is the author of the blog
type Author struct {
	Username string `json:"username,omitempty" bson:"username,omitempty"`
	Email    string `json:"email,omitempty" bson:"email,omitempty"`
}

// TODO: Implement Blog and Author validation
