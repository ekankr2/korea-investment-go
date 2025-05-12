package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"korea-investment-go/lib"
	"korea-investment-go/routes"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found")
	}
}

func main() {
	config := GetConfig()

	lib.SetupRedis(
		config.RedisAddr,
		config.RedisPass,
		0,
	)

	gin.SetMode(config.GinMode)

	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":" + config.Port)
}
