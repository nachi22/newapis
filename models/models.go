package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
	Email    string             `json:"email" bson:"email,omitempty"`
	Password string             `json:"password" bson:"password,omitempty"`
}

type Posts struct {
	ID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Caption         string             `json:"caption,omitempty" bson:"caption,omitempty"`
	ImageURL        string             `json:"imageurl,omitempty" bson:"imageurl,omitempty"`
	PostedTimestamp string             `json:"postedtimestamp" bson:"postedtimestamp,omitempty"`
	User            *User              `json:"user" bson:"user,omitempty"`
}
