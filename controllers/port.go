package handlers

import (
	"context"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"vectorcd/models"
	"vectorcd/connect"
)

func GetUsedPorts(collection *mongo.Collection) ([]string, error) {
	var apps []models.App
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

		if !isUsed {
			return randomPort, nil
		}
	}
}

func GetRandomPort(c *fiber.Ctx) error {
	client := mongodb.ConnectDB()
	collection := mongodb.GetCollection(client, "port") // Replace with your actual collection name

	minPort := 8000
	maxPort := 9000

	randomPort, err := GetUnusedPort(collection, minPort, maxPort)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get an unused port",
		})
	}

	return c.JSON(fiber.Map{
		"unused_port": randomPort,
	})
}
