package routes

import (
	"github.com/gin-gonic/gin"
	"korea-investment-go/services"
	"net/http"
	"os"
)

type KISHandler struct {
	kisService *services.KISService
}

func NewKISHandler() *KISHandler {
	// 서비스 생성
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

// init 함수에서 라우트 등록
func init() {
	RegisterRoutes(KISRoutes)
}

// KISRoutes 함수는 한국투자증권 API 관련 라우트를 설정합니다
func KISRoutes(r *gin.Engine) {
	// 핸들러 생성
	handler := NewKISHandler()

	// API 라우트 설정
	api := r.Group("/api/kis")
	{
		api.GET("/stock/:stockCode", handler.GetStockPrice)
		api.GET("/account/:accountNo", handler.GetAccountBalance)
	}
}
