package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"wallet-app/internal/api/handlers"
	"wallet-app/internal/models"
	"wallet-app/internal/service"
	"wallet-app/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetTransactionsHandler(t *testing.T) {
	_, transactionRepo, walletRepo, userRepo := utils.SetupTestDB()
	walletService := service.NewWalletServiceImpl(transactionRepo, walletRepo)
	walletHandler := handlers.NewWalletHandler(walletService)
	r := gin.Default()
	r.GET("/wallet/:id/transactions", walletHandler.GetTransactionsHandler)

	userID, _ := userRepo.Save(models.User{Name: "Test User"})

	walletID, _ := walletRepo.Save(models.Wallet{Balance: 100.0, UserID: userID})

	transactionRepo.Save(models.Transaction{WalletID: walletID, Amount: 50.0, Type: "deposit"})
	transactionRepo.Save(models.Transaction{WalletID: walletID, Amount: 20.0, Type: "withdrawal"})

	tests := []struct {
		name           string
		walletID       string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid Wallet ID",
			walletID:       strconv.Itoa(int(walletID)),
			expectedStatus: http.StatusOK,
			expectedBody:   "transactions",
		},
		{
			name:           "Invalid Wallet ID",
			walletID:       "invalid-id",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"invalid wallet id"}`,
		},
		{
			name:           "Non-existent Wallet",
			walletID:       "999",
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"error":"resource not found"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest(http.MethodGet, "/wallet/"+tt.walletID+"/transactions", nil)
			resp := httptest.NewRecorder()

			r.ServeHTTP(resp, req)

			assert.Equal(t, tt.expectedStatus, resp.Code)
			assert.Contains(t, resp.Body.String(), tt.expectedBody)
		})
	}
}
