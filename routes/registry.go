package routes

import "github.com/gin-gonic/gin"

type RouteRegistrar func(*gin.Engine)

var registrars []RouteRegistrar

func RegisterRoutes(registrar RouteRegistrar) {
	registrars = append(registrars, registrar)
}

func RegisterAllRoutes() {
	RegisterRoutes(KISRoutes)
}

func SetupRoutes(r *gin.Engine) {
	for _, registrar := range registrars {
		registrar(r)
	}
}

func init() {
	RegisterAllRoutes()
}
