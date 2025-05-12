package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	AppKey     string
	AppSecret  string
	GinMode    string
	Port       string
	RedisAddr  string
	RedisPass  string
	AccountNum string
}

var cfg *Config

// 설정을 초기화하는 함수 (init 대신 명시적으로 호출)
func InitConfig() *Config {
	// .env 파일 로드
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found")
	}

	// 환경 변수 로드
	cfg = &Config{
		AppKey:     os.Getenv("APP_KEY"),
		AppSecret:  os.Getenv("APP_SECRET"),
		GinMode:    getEnvWithDefault("GIN_MODE", "debug"),
		Port:       getEnvWithDefault("PORT", "8080"),
		RedisAddr:  os.Getenv("REDIS_ADDRESS"),
		RedisPass:  os.Getenv("REDIS_PASSWORD"),
		AccountNum: os.Getenv("ACCOUNT_NUMBER"),
	}

	return cfg
}

func getEnvWithDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func GetConfig() *Config {
	if cfg == nil {
		return InitConfig()
	}
	return cfg
}
