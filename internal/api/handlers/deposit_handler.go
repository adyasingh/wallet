package handlers

import (
	"net/http"
	"strconv"
	"wallet-app/internal/errors"
	"wallet-app/internal/service"

	"github.com/gin-gonic/gin"
)

func (c *WalletHandler) DepositHandler(ctx *gin.Context) {
	walletID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidWalletID.Error()})
		return
	}
	req := service.CreateDepositRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.WalletID = uint(walletID)
	err = c.walletService.Deposit(req)
	if err != nil {
		status, message := errors.HttpError(err)
		ctx.JSON(status, gin.H{"error": message})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Deposit successful"})
}
