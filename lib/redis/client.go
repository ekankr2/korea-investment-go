package redis

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

func Setup(addr, password string, db int) {
	Client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	_, err := Client.Ping(Ctx).Result()
	if err != nil {
		log.Printf("Failed to connect to Redis: %v", err)
	} else {
		log.Println("Connected to Redis successfully")
	}
}

func Get(key string) (string, error) {
	val, err := Client.Get(Ctx, key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("key %s does not exist", key)
	} else if err != nil {
		return "", err
	}
	return val, nil
}

func Set(key string, value interface{}, expiration time.Duration) error {
	log.Printf("Redis.Set 함수 호출: key=%s, expiration=%v", key, expiration)
	err := Client.Set(Ctx, key, value, expiration).Err()
	if err != nil {
		log.Printf("Redis.Set 실패: %v", err)
		return err
	}
	log.Printf("Redis.Set 성공")
	return nil
}

func Delete(key string) error {
	return Client.Del(Ctx, key).Err()
}

func HashSet(key string, values map[string]interface{}) error {
	return Client.HSet(Ctx, key, values).Err()
}

func HashGet(key, field string) (string, error) {
	val, err := Client.HGet(Ctx, key, field).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("key %s or field %s does not exist", key, field)
	} else if err != nil {
		return "", err
	}
	return val, nil
}

func HashGetAll(key string) (map[string]string, error) {
	return Client.HGetAll(Ctx, key).Result()
}
