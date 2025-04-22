package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	if value, ok := os.LookupEnv("MONGO_URI"); ok {
		return value
	}
	return "mongodb+srv://anumoluakash2004:hd9bHTUzsvYtlfxt@apscsci.w4ny8bh.mongodb.net/"
}
func GetAPIKey() string {
	// Load environment variables from .env file
	err := godotenv.Load(".env") // Ensure the correct file name
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get the API key from environment variables
	apiKey := os.Getenv("APIkey")
	if apiKey == "" {
		log.Println("Warning: API key is empty or not found in .env")
	}

	return apiKey
}
