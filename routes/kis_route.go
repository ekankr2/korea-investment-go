package routes

import (
	"github.com/gin-gonic/gin"
	"korea-investment-go/services"
	"net/http"
)

type KisRouter struct {
	kisService *services.KISService
}

func NewKISRouter(kisService *services.KISService) *KisRouter {
	return &KisRouter{
		kisService: kisService,
	}
}

func (c *KisRouter) GetStockPrice(ctx *gin.Context) {
	stockCode := ctx.Param("stockCode")
	if stockCode == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Stock code is required"})
		return
	}

	result, err := c.kisService.GetStockPrice(stockCode)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *KisRouter) GetAccountBalance(ctx *gin.Context) {
	accountNo := ctx.Param("accountNo")
	if accountNo == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Account number is required"})
		return
	}

	result, err := c.kisService.GetAccountBalance(accountNo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
