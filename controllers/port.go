package controllers

import (
	"context"
	"math/rand"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	
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

func GetUsedPorts(collection *mongo.Collection) ([]string, error) {
	var apps []App
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.Background(), &apps); err != nil {
		return nil, err
	}

	var usedPorts []string
	for _, app := range apps {
		usedPorts = append(usedPorts, app.Port)
	}

	return usedPorts, nil
}

func GenerateRandomPort(minPort, maxPort int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(maxPort-minPort+1) + minPort
}

// GetUnusedPort finds an unused port by checking against the list of used ports
func GetUnusedPort(collection *mongo.Collection, minPort, maxPort int) (int, error) {
	usedPorts, err := GetUsedPorts(collection)
	if err != nil {
		return 0, err
	}

	for {
		randomPort := GenerateRandomPort(minPort, maxPort)
		portStr := strconv.Itoa(randomPort)

		isUsed := false
		for _, usedPort := range usedPorts {
			if portStr == usedPort {
				isUsed = true
				break
			}
		}

		// If the port is not in the used ports, return it
		if !isUsed {
			return randomPort, nil
		}
	}
}
