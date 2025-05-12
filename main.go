package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"korea-investment-go/config"
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
	initConfig := config.InitConfig()

	redis.Setup(
		initConfig.RedisAddr,
		initConfig.RedisPass,
		0,
	)

	gin.SetMode(initConfig.GinMode)

	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":" + initConfig.Port)
}
