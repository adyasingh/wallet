package router_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"wallet-app/internal/api/handlers"
	"wallet-app/internal/models"
	"wallet-app/internal/router"
	"wallet-app/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type MockWalletService struct{}

func (m MockWalletService) Deposit(req service.CreateDepositRequest) error {
	return nil
}

func (m MockWalletService) GetBalance(id uint) (float64, error) {
	return 100.0, nil
}

func (m MockWalletService) GetTransactions(id uint) ([]models.Transaction, error) {
	return []models.Transaction{{Amount: 50.0}}, nil
}

func (m MockWalletService) Withdraw(req service.CreateWithdrawalRequest) error {
	return nil
}

func (m MockWalletService) Transfer(req service.TransferRequest) error {
	return nil
}

func TestRouter(t *testing.T) {
	gin.SetMode(gin.TestMode)
	walletHandler := handlers.NewWalletHandler(MockWalletService{})
	router := router.NewRouter(walletHandler)

	tests := []struct {
		method             string
		path               string
		body               interface{}
		expectedStatusCode int
		expectedBody       gin.H
	}{
		{
			method:             "GET",
			path:               "/api/wallet/1/balance",
			body:               nil,
			expectedStatusCode: http.StatusOK,
			expectedBody:       gin.H{"balance": 100.0},
		},
		{
			method:             "GET",
			path:               "/api/wallet/1/transactions",
			body:               nil,
			expectedStatusCode: http.StatusOK,
			expectedBody:       gin.H{"transactions": []models.Transaction{{Amount: 50.0}}},
		},
		{
			method:             "POST",
			path:               "/api/wallet/1/transfer",
			body:               service.TransferRequest{Amount: 50, ToWalletId: 2},
			expectedStatusCode: http.StatusOK,
			expectedBody:       gin.H{"message": "Transfer successful"},
		},
		{
			method:             "POST",
			path:               "/api/wallet/1/deposit",
			body:               service.CreateDepositRequest{Amount: 50},
			expectedStatusCode: http.StatusOK,
			expectedBody:       gin.H{"message": "Deposit successful"},
		},
		{
			method:             "POST",
			path:               "/api/wallet/1/withdraw",
			body:               service.CreateWithdrawalRequest{Amount: 50},
			expectedStatusCode: http.StatusOK,
			expectedBody:       gin.H{"message": "Withdrawal successful"},
		},
		{
			method:             "GET",
			path:               "/api/wallet/unknown",
			body:               nil,
			expectedStatusCode: http.StatusNotFound,
			expectedBody:       gin.H{"message": "Page not found"},
		},
	}
	for _, test := range tests {
		var reqBody []byte
		if test.body != nil {
			var err error
			reqBody, err = json.Marshal(test.body)
			assert.NoError(t, err)
		}

		req := httptest.NewRequest(test.method, test.path, bytes.NewBuffer(reqBody))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, test.expectedStatusCode, w.Code)

		expectedBodyJSON, err := json.Marshal(test.expectedBody)
		assert.NoError(t, err)
		assert.JSONEq(t, string(expectedBodyJSON), w.Body.String())
	}
}
