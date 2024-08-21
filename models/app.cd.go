package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type App struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`  // MongoDB ObjectID
	AppName   string             `bson:"first_name" json:"first_name"`       // User's first name
	Port      string             `bson:"last_name" json:"last_name"`         // User's last name
	Email     string             `bson:"email" json:"email"`                 // User's email address
	GithubId  string             `bson:"password" json:"-"`                  // User's password (hash it before storing)
	CreatedAt primitive.DateTime `bson:"created_at,omitempty" json:"created_at,omitempty"` // Timestamp for when the user was created
	UpdatedAt primitive.DateTime `bson:"updated_at,omitempty" json:"updated_at,omitempty"` // Timestamp for when the user was last updated
}
