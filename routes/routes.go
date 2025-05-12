package routes

import (
	"github.com/gin-gonic/gin"
	"korea-investment-go/services"
)

type RouteRegistrar func(*gin.Engine)

var registrars []RouteRegistrar
var kisService *services.KISService

func RegisterRoutes(registrar RouteRegistrar) {
	registrars = append(registrars, registrar)
}

func RegisterAllRoutes() {
	RegisterRoutes(AccountRoutes)
	RegisterRoutes(StockRoutes)
	RegisterRoutes(OrderRoutes)
}

func SetupRoutes(r *gin.Engine) {
	kisService = services.NewKISService()

	for _, registrar := range registrars {
		registrar(r)
	}

}

func init() {
	RegisterAllRoutes()
}
