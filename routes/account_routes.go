package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AccountRoutes(r *gin.Engine) {
	accountGroup := r.Group("/kis/account")
	{
		accountGroup.GET("/:accountNo", GetAccountBalance)
	}
}

// GetAccountBalance 핸들러는 계좌 잔고를 조회합니다
func GetAccountBalance(ctx *gin.Context) {
	accountNo := ctx.Param("accountNo")
	if accountNo == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Account number is required"})
		return
	}

	result, err := kisService.GetAccountBalance(accountNo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
