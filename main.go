package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// Gin 엔진 생성
	r := gin.Default()

	// 라우트 설정
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "안녕하세요! Gin 웹 프레임워크에 오신 것을 환영합니다!",
		})
	})

	// 서버 실행
	r.Run(":8080") // 기본 포트는 8080
}
