package handlers_test

import (
	"bytes"
	"encoding/json"
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

func TestWithdrawalHandler(t *testing.T) {
	_, transactionRepo, walletRepo, userRepo := utils.SetupTestDB()
	walletService := service.NewWalletServiceImpl(transactionRepo, walletRepo)
	walletHandler := handlers.NewWalletHandler(walletService)
	r := gin.Default()
	r.POST("/wallet/:id/withdraw", walletHandler.WithdrawalHandler)

	userID, _ := userRepo.Save(models.User{Name: "Test User"})
	walletID, _ := walletRepo.Save(models.Wallet{Balance: 100.0, UserID: userID})

	tests := []struct {
		name             string
		walletID         string
		withdrawalAmount float64
		expectedStatus   int
		expectedBody     string
	}{
		{
			name:             "Valid Withdrawal",
			walletID:         strconv.Itoa(int(walletID)),
			withdrawalAmount: 50.0,
			expectedStatus:   http.StatusOK,
			expectedBody:     `{"message":"Withdrawal successful"}`,
		},
		{
			name:             "Invalid Wallet ID",
			walletID:         "invalid-id",
			withdrawalAmount: 50.0,
			expectedStatus:   http.StatusBadRequest,
			expectedBody:     `{"error":"invalid wallet id"}`,
		},
		{
			name:             "Invalid Withdrawal Amount",
			walletID:         strconv.Itoa(int(walletID)),
			withdrawalAmount: -50.0,
			expectedStatus:   http.StatusBadRequest,
			expectedBody:     `{"error":"amount must be greater than zero"}`,
		},
		{
			name:             "Insufficient Funds",
			walletID:         strconv.Itoa(int(walletID)),
			withdrawalAmount: 150.0,
			expectedStatus:   http.StatusBadRequest,
			expectedBody:     `{"error":"insufficient funds"}`,
		},
		{
			name:             "Non-existent Wallet",
			walletID:         "999",
			withdrawalAmount: 50.0,
			expectedStatus:   http.StatusNotFound,
			expectedBody:     `{"error":"resource not found"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := gin.H{"amount": tt.withdrawalAmount}
			jsonBody, _ := json.Marshal(body)
			req, _ := http.NewRequest(http.MethodPost, "/wallet/"+tt.walletID+"/withdraw", bytes.NewBuffer(jsonBody))
			resp := httptest.NewRecorder()

			r.ServeHTTP(resp, req)

			assert.Equal(t, tt.expectedStatus, resp.Code)
			assert.Contains(t, resp.Body.String(), tt.expectedBody)
		})
	}
}
