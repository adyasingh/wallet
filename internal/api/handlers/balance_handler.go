package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"wallet-app/internal/errors"

	"github.com/gin-gonic/gin"
)

func (c *WalletHandler) GetBalanceHandler(ctx *gin.Context) {
	walletID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid wallet id"})
		return
	}
	balance, err := c.walletService.GetBalance(uint(walletID))
	if err != nil {
		status, message := errors.HttpError(err)
		ctx.JSON(status, gin.H{"error": message})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"balance": balance})
}
