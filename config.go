package main

import "os"

type Config struct {
	AppKey    string
	AppSecret string
	GinMode   string
	Port      string
}

var cfg *Config

func init() {
	// 환경 변수 로드
	cfg = &Config{
		AppKey:    os.Getenv("APP_KEY"),
		AppSecret: os.Getenv("APP_SECRET"),
		GinMode:   getEnvWithDefault("GIN_MODE", "debug"),
		Port:      getEnvWithDefault("PORT", "8080"),
	}
}

func getEnvWithDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func GetConfig() *Config {
	return cfg
}
