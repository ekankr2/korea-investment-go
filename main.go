package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	config2 "korea-investment-go/config"
	"korea-investment-go/lib/redis"
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
	config := config2.InitConfig()

	redis.Setup(
		config.RedisAddr,
		config.RedisPass,
		0,
	)

	gin.SetMode(config.GinMode)

	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":" + config.Port)
}
