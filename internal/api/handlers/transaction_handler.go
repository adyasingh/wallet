package handlers

import (
	"net/http"
	"strconv"
	"wallet-app/internal/errors"

	"github.com/gin-gonic/gin"
)

func (c *WalletHandler) GetTransactionsHandler(ctx *gin.Context) {
	walletID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidWalletID.Error()})
		return
	}

	transactions, err := c.walletService.GetTransactions(uint(walletID))
	if err != nil {
		status, message := errors.HttpError(err)
		ctx.JSON(status, gin.H{"error": message})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"transactions": transactions})
}
