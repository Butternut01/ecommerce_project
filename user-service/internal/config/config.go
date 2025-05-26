package config

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	ServerPort   string
	DatabaseURI  string
	DatabaseName string
	EmailFrom    string
    EmailPass    string
    SMTPHost     string
    SMTPPort     string
}

func NewConfig() *Config {
	return &Config{
		ServerPort:   getEnv("USER_SERVICE_PORT", "50051"),
		DatabaseURI:  getEnv("USER_SERVICE_DB_URI", "mongodb://localhost:27017"),
		DatabaseName: getEnv("USER_SERVICE_DB_NAME", "userdb"),
		EmailFrom:    getEnv("EMAIL_FROM", "asanalibayadilov555@gmail.com"),
        EmailPass:    getEnv("EMAIL_PASS", "nigwfqchzouwwttz"), // Use app password
        SMTPHost:     getEnv("SMTP_HOST", "smtp.gmail.com"),
        SMTPPort:     getEnv("SMTP_PORT", "587"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func ConnectDatabase(uri string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	// Ping the database to verify the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}