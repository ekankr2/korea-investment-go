package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func StockRoutes(r *gin.Engine) {
	stockGroup := r.Group("/kis/stock")
	{
		stockGroup.GET("/:stockCode", GetStockPrice)
		stockGroup.GET("/daily/:stockCode", GetDailyStockPrice)
	}
}

// 주식현재가 시세[v1_국내주식-008]
func GetStockPrice(ctx *gin.Context) {
	stockCode := ctx.Param("stockCode")
	if stockCode == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Stock code is required"})
		return
	}

	result, err := kisService.GetStockPrice(stockCode)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

// 주식현재가 일자별[v1_국내주식-010]
func GetDailyStockPrice(ctx *gin.Context) {
	stockCode := ctx.Param("stockCode")
	if stockCode == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Stock code is required"})
		return
	}

	startDate := ctx.Query("startDate")
	endDate := ctx.Query("endDate")
	periodCode := ctx.DefaultQuery("period", "D") // 기본값은 일봉(D)

	// 필수 파라미터 검증
	if startDate == "" || endDate == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "startDate and endDate are required"})
		return
	}

	// periodCode 검증
	if periodCode != "D" && periodCode != "W" && periodCode != "M" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "period must be one of D(daily), W(weekly), or M(monthly)"})
		return
	}

	result, err := kisService.GetDailyStockPrice(stockCode, startDate, endDate, periodCode)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
