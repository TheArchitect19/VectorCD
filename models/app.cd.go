package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type App struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`  
	AppName   string             `bson:"app_name" json:"app_name"`      
	Port      string             `bson:"port" json:"port"`       
	Email     *string             `bson:"email" json:"email"`           
	GithubId  *string             `bson:"github" json:"github"`               
	CreatedAt primitive.DateTime `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt primitive.DateTime `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}
