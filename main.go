package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

var ctx = context.Background()

type Person struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Occupation string `json:"occupation"`
}

func main() {
	fmt.Println("Hello, World!")

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// Ping Redis server
	if _, err := client.Ping(ctx).Result(); err != nil {
		fmt.Println("Error connecting to Redis:", err)
		return
	}
	fmt.Println("Connected to Redis successfully")

	// Create a new person
	eID := uuid.NewString()
	person := Person{
		ID:         eID,
		Name:       "Moneypenny",
		Age:        24,
		Occupation: "Programmer",
	}

	// Serialize person to JSON
	jsonString, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Store JSON in Redis
	err = client.Set(ctx, eID, jsonString, 0).Err()
	if err != nil {
		fmt.Println("Error saving to Redis:", err)
		return
	}

	fmt.Println("Person saved successfully with ID:", eID)

	// Retrieve JSON from Redis
	val, err := client.Get(ctx, eID).Result()
	if err != nil {
		fmt.Println("Error retrieving from Redis:", err)
		return
	}

	fmt.Println("Data retrieved from Redis:", val)

	// Deserialize JSON back to struct
	var retrievedPerson Person
	if err := json.Unmarshal([]byte(val), &retrievedPerson); err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	fmt.Printf("Deserialized Person: %+v\n", retrievedPerson)
}