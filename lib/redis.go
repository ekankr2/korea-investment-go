package lib

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	Client *redis.Client
	Ctx    = context.Background()
)

// Setup initializes the Redis client
func SetupRedis(addr, password string, db int) {
	Client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	// Test the connection
	_, err := Client.Ping(Ctx).Result()
	if err != nil {
		log.Printf("Failed to connect to Redis: %v", err)
	} else {
		log.Println("Connected to Redis successfully")
	}
}

// Get retrieves a value from Redis
func Get(key string) (string, error) {
	val, err := Client.Get(Ctx, key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("key %s does not exist", key)
	} else if err != nil {
		return "", err
	}
	return val, nil
}

// Set stores a value in Redis with optional expiration
func Set(key string, value interface{}, expiration time.Duration) error {
	return Client.Set(Ctx, key, value, expiration).Err()
}

// Delete removes a key from Redis
func Delete(key string) error {
	return Client.Del(Ctx, key).Err()
}

// HashSet sets multiple fields in a Redis hash
func HashSet(key string, values map[string]interface{}) error {
	return Client.HSet(Ctx, key, values).Err()
}

// HashGet retrieves a field from a Redis hash
func HashGet(key, field string) (string, error) {
	val, err := Client.HGet(Ctx, key, field).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("key %s or field %s does not exist", key, field)
	} else if err != nil {
		return "", err
	}
	return val, nil
}

// HashGetAll retrieves all fields from a Redis hash
func HashGetAll(key string) (map[string]string, error) {
	return Client.HGetAll(Ctx, key).Result()
}
