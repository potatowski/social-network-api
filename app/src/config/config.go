package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// DatabaseUrl is the url to connect in database
	DatabaseUrl = ""

	// Port where API will be running
	Port = 0

	// SecretKey is the key to generate tokens
	SecretKey []byte
)

// Load will initialize environment variables
func Load() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 5000
	}

	DatabaseUrl = os.Getenv("DATABASE_URL")
	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
