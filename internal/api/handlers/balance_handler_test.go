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

func TestGetBalanceHandler(t *testing.T) {
	_, transactionRepo, walletRepo, userRepo := utils.SetupTestDB()

	walletService := service.NewWalletServiceImpl(transactionRepo, walletRepo)
	walletHandler := handlers.NewWalletHandler(walletService)

	r := gin.Default()
	r.GET("/wallet/:id/balance", walletHandler.GetBalanceHandler)

	userID, _ := userRepo.Save(models.User{Name: "Test User"})
	walletID, _ := walletRepo.Save(models.Wallet{Balance: 100.0, UserID: userID})

	tests := []struct {
		name           string
		walletID       string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid Request",
			walletID:       strconv.Itoa(int(walletID)),
			expectedStatus: http.StatusOK,
			expectedBody:   `{"balance":100}`,
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
			req, _ := http.NewRequest(http.MethodGet, "/wallet/"+tt.walletID+"/balance", nil)
			resp := httptest.NewRecorder()

			r.ServeHTTP(resp, req)

			assert.Equal(t, tt.expectedStatus, resp.Code)
			assert.JSONEq(t, tt.expectedBody, resp.Body.String())
		})
	}
}
