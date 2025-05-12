package routes

import (
	"github.com/gin-gonic/gin"
	"korea-investment-go/services"
	"os"
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
}

func SetupRoutes(r *gin.Engine) {
	kisService = services.NewKISService(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
	)

	for _, registrar := range registrars {
		registrar(r)
	}

}

func init() {
	RegisterAllRoutes()
}
