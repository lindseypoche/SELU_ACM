package users

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User struct ...
type User struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName   string             `json:"first_name,omitempty" bson:"firstname,omitempty"`
	LastName    string             `json:"last_name,omitempty" bson:"last_name,omitempty"`
	Email       string             `json:"email,omitempty" bson:"email,omitempty"`
	Password    string             `json:"-" bson:"-"`
	DateCreated string             `json:"date_created"`
}

// TODO: implement user validation
