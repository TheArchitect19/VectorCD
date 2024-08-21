package handlers

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"vectorcd/connect"
	"vectorcd/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"vectorcd/config"
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
	nginxConfig := fmt.Sprintf(`
server {
    listen 80;
    server_name api-vl.gokapturehub.com api.gokapturehub.com;

    location / {
        proxy_pass http://127.0.0.1:%d/;
        proxy_redirect off;

        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;

        client_max_body_size 50m;
        client_body_buffer_size 128k;

        proxy_connect_timeout 90;
        proxy_send_timeout 90;
        proxy_read_timeout 90;

        proxy_buffer_size 4k;
        proxy_buffers 4 32k;
        proxy_busy_buffers_size 64k;
        proxy_temp_file_write_size 64k;
    }
}
`, randomPort)

	// Define the path where the file should be saved
	filePath := "/etc/nginx/sites-enabled/gokapturehub.conf"

	// Create or overwrite the file at the specified path
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)

	}
	defer file.Close()

	// Write the configuration content to the file
	_, err = file.WriteString(nginxConfig)
	if err != nil {
		fmt.Println("Error writing to file:", err)

	}
	config.ReloadNginx();
	fmt.Println("NGINX configuration file created successfully at", filePath)

	return c.JSON(fiber.Map{
		"unused_port": randomPort,
	})
}
