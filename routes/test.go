package routes

import "github.com/gin-gonic/gin"

func init() {
	RegisterRoutes(TestRoutes)
}

func TestRoutes(r *gin.Engine) {
	group := r.Group("/users")
	group.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Get all users",
		})
	})
}

// 핸들러 함수들...
