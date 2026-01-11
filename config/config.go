package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Version      string
	ServiceName  string
	HttpPort     int64
	JwtSecretKey string
}

var config *Config

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load the .env file:", err)
		os.Exit(1)
	}

	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("Please provide a valid version")
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if version == "" {
		fmt.Println("Please provide a valid service name")
		os.Exit(1)
	}

	httpPort := os.Getenv("HTTP_PORT")
	if version == "" {
		fmt.Println("Please provide a valid HTTP port")
		os.Exit(1)
	}

	port, err := strconv.ParseInt(httpPort, 10, 64)
	if err != nil {
		fmt.Println("Failed to convert the HTTP port provided into Integer value:", err)
		os.Exit(1)
	}

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if version == "" {
		fmt.Println("Please provide a valid JWT Secret KEy")
		os.Exit(1)
	}

	config = &Config{
		Version:      version,
		ServiceName:  serviceName,
		HttpPort:     port,
		JwtSecretKey: jwtSecretKey,
	}
}

func GetConfig() *Config {
	if config == nil {
		loadConfig()
	}

	return config
}
