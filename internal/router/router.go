package router

import (
	"wallet-app/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func NewRouter(walletHandler *handlers.WalletHandler) *gin.Engine {
	service := gin.Default()

	service.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Page not found"})
	})

	router := service.Group("/api")
	walletRouter := router.Group("/wallet")
	walletRouter.GET("/:id/balance", walletHandler.GetBalanceHandler)
	walletRouter.GET("/:id/transactions", walletHandler.GetTransactionsHandler)
	walletRouter.POST("/:id/transfer", walletHandler.TransferHandler)
	walletRouter.POST("/:id/deposit", walletHandler.DepositHandler)
	walletRouter.POST("/:id/withdraw", walletHandler.WithdrawalHandler)

	return service
}
