package handlers

import (
	"net/http"
	"strconv"
	"wallet-app/internal/errors"
	"wallet-app/internal/service"

	"github.com/gin-gonic/gin"
)

func (c *WalletHandler) TransferHandler(ctx *gin.Context) {
	walletID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidWalletID.Error()})
		return
	}
	req := service.TransferRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.FromWalletId = uint(walletID)
	err = c.walletService.Transfer(req)
	if err != nil {
		status, message := errors.HttpError(err)
		ctx.JSON(status, gin.H{"error": message})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Transfer successful"})
}
