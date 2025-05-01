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

func TestDepositHandler(t *testing.T) {
	_, transactionRepo, walletRepo, userRepo := utils.SetupTestDB()

	walletService := service.NewWalletServiceImpl(transactionRepo, walletRepo)
	walletHandler := handlers.NewWalletHandler(walletService)

	r := gin.Default()
	r.POST("/wallet/:id/deposit", walletHandler.DepositHandler)

	userID, err := userRepo.Save(models.User{Name: "Test User"})
	assert.NoError(t, err)

	walletID, err := walletRepo.Save(models.Wallet{Balance: 100.0, UserID: userID})
	assert.NoError(t, err)

	tests := []struct {
		name           string
		walletID       string
		depositAmount  float64
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid Deposit",
			walletID:       strconv.Itoa(int(walletID)),
			depositAmount:  50.0,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"message":"Deposit successful"}`,
		},
		{
			name:           "Invalid Wallet ID",
			walletID:       "invalid-id",
			depositAmount:  50.0,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"invalid wallet id"}`,
		},
		{
			name:           "Invalid Deposit Amount",
			walletID:       strconv.Itoa(int(walletID)),
			depositAmount:  -50.0,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"amount must be greater than zero"}`,
		},
		{
			name:           "Non-existent Wallet",
			walletID:       "999",
			depositAmount:  50.0,
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"error":"resource not found"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := gin.H{"amount": tt.depositAmount}
			jsonBody, _ := json.Marshal(body)
			req, _ := http.NewRequest(http.MethodPost, "/wallet/"+tt.walletID+"/deposit", bytes.NewBuffer(jsonBody))
			resp := httptest.NewRecorder()

			r.ServeHTTP(resp, req)

			assert.Equal(t, tt.expectedStatus, resp.Code)
			assert.JSONEq(t, tt.expectedBody, resp.Body.String())
		})
	}
}
