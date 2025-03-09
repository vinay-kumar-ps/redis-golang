package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	// Initialize Redis client
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// Check Redis connection
	if _, err := client.Ping(ctx).Result(); err != nil {
		fmt.Println("Error connecting to Redis:", err)
		return
	}
	fmt.Println("Connected to Redis!")

	// Simulate URL Shortening
	longURL := "https://example.com/user/details/434345"
	shortURL := generateShortURL()

	// Store the mapping (shortURL -> longURL)
	err := client.Set(ctx, shortURL, longURL, 24*time.Hour).Err() // Expires in 24 hours
	if err != nil {
		fmt.Println("Error saving URL:", err)
		return
	}
	fmt.Printf("Original URL: %s\nShort URL: http://short.ly/%s\n", longURL, shortURL)

	// Retrieve original URL
	retrievedURL, err := client.Get(ctx, shortURL).Result()
	if err != nil {
		fmt.Println("Error retrieving URL:", err)
		return
	}
	fmt.Println("Retrieved URL:", retrievedURL)
}

// Helper function to generate a random short URL
func generateShortURL() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())

	short := make([]byte, 6)
	for i := range short {
		short[i] = charset[rand.Intn(len(charset))]
	}
	return string(short)
}
