package routes

import (
	"github.com/gin-gonic/gin"
	"korea-investment-go/services"
	"net/http"
)

func OrderRoutes(r *gin.Engine) {
	orderGroup := r.Group("/kis/order")
	{
		orderGroup.POST("/cash", PostOrderCash)
	}
}

// 주식주문(현금)[v1_국내주식-001]
func PostOrderCash(ctx *gin.Context) {
	var req services.OrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	orderType := ctx.DefaultQuery("type", "buy") // 쿼리로 buy/sell 구분

	result, err := kisService.OrderCash(req, orderType)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
