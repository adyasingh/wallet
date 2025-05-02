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

func TestTransferHandler(t *testing.T) {
	_, transactionRepo, walletRepo, userRepo := utils.SetupTestDB()
	walletService := service.NewWalletServiceImpl(transactionRepo, walletRepo)
	walletHandler := handlers.NewWalletHandler(walletService)

	r := gin.Default()
	r.POST("/wallet/:id/transfer", walletHandler.TransferHandler)

	userID1, err := userRepo.Save(models.User{Name: "Test User"})
	assert.NoError(t, err)
	userID2, err := userRepo.Save(models.User{Name: "Test User2"})
	assert.NoError(t, err)

	walletID1, err := walletRepo.Save(models.Wallet{Balance: 100.0, UserID: userID1})
	assert.NoError(t, err)
	walletID2, err := walletRepo.Save(models.Wallet{Balance: 100.0, UserID: userID2})
	assert.NoError(t, err)

	tests := []struct {
		name           string
		walletID       string
		amount         float64
		ToWalletId     uint
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid Transfer",
			walletID:       strconv.Itoa(int(walletID1)),
			amount:         50.0,
			ToWalletId:     walletID2,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"message":"Transfer successful"}`,
		},
		{
			name:           "Invalid Wallet ID",
			walletID:       "invalid-id",
			amount:         50.0,
			ToWalletId:     walletID2,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"invalid wallet id"}`,
		},
		{
			name:           "Invalid Amount",
			walletID:       strconv.Itoa(int(walletID1)),
			amount:         -50.0,
			ToWalletId:     walletID2,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"amount must be greater than zero"}`,
		},
		{
			name:           "Non-existent outgoing wallet",
			walletID:       "999",
			amount:         50.0,
			ToWalletId:     walletID2,
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"error":"resource not found"}`,
		},
		{
			name:           "Non-existent incoming wallet",
			walletID:       strconv.Itoa(int(walletID1)),
			ToWalletId:     999,
			amount:         50.0,
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"error":"resource not found"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := gin.H{"amount": tt.amount, "to_wallet_id": tt.ToWalletId}
			jsonBody, err := json.Marshal(body)
			assert.NoError(t, err)
			req, err := http.NewRequest(http.MethodPost, "/wallet/"+tt.walletID+"/transfer", bytes.NewBuffer(jsonBody))
			assert.NoError(t, err)
			resp := httptest.NewRecorder()

			r.ServeHTTP(resp, req)

			assert.Equal(t, tt.expectedStatus, resp.Code)
			assert.JSONEq(t, tt.expectedBody, resp.Body.String())
		})
	}
}
