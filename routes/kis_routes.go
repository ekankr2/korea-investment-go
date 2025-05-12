package routes

import (
	"github.com/gin-gonic/gin"
	"korea-investment-go/services"
	"net/http"
	"os"
)

func KISRoutes(r *gin.Engine) {
	handler := NewKISHandler()

	api := r.Group("/kis")
	{
		api.GET("/stock/:stockCode", handler.GetStockPrice)
		api.GET("/account/:accountNo", handler.GetAccountBalance)
		api.GET("/stock/daily/:stockCode", handler.GetDailyStockPrice)
	}
}

type KISHandler struct {
	kisService *services.KISService
}

func NewKISHandler() *KISHandler {
	kisService := services.NewKISService(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
	)

	return &KISHandler{
		kisService: kisService,
	}
}

// GetStockPrice 핸들러는 주식 시세를 조회합니다
func (h *KISHandler) GetStockPrice(ctx *gin.Context) {
	stockCode := ctx.Param("stockCode")
	if stockCode == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Stock code is required"})
		return
	}

	result, err := h.kisService.GetStockPrice(stockCode)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

// GetAccountBalance 핸들러는 계좌 잔고를 조회합니다
func (h *KISHandler) GetAccountBalance(ctx *gin.Context) {
	accountNo := ctx.Param("accountNo")
	if accountNo == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Account number is required"})
		return
	}

	result, err := h.kisService.GetAccountBalance(accountNo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

// GetDailyStockPrice 핸들러는 주식의 일/주/월 시세를 조회합니다
func (h *KISHandler) GetDailyStockPrice(ctx *gin.Context) {
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

	result, err := h.kisService.GetDailyStockPrice(stockCode, startDate, endDate, periodCode)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
