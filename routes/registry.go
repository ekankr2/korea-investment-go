package routes

import "github.com/gin-gonic/gin"

type RouteRegistrar func(*gin.Engine)

var registrars []RouteRegistrar

// RegisterRoutes 함수는 라우트 등록 함수를 레지스트리에 추가합니다
func RegisterRoutes(registrar RouteRegistrar) {
	registrars = append(registrars, registrar)
}

// SetupRoutes 함수는 등록된 모든 라우트를 엔진에 설정합니다
func SetupRoutes(r *gin.Engine) {
	for _, registrar := range registrars {
		registrar(r)
	}
}
