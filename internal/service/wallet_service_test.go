package service_test

import (
	"testing"
	"wallet-app/internal/errors"
	"wallet-app/internal/models"
	"wallet-app/internal/service"
	"wallet-app/internal/utils"

	"github.com/stretchr/testify/assert"
)

func TestWalletService_Deposit(t *testing.T) {
	_, transactionRepo, walletRepo, userRepo := utils.SetupTestDB()

	walletService := service.NewWalletServiceImpl(transactionRepo, walletRepo)
	userID, _ := userRepo.Save(models.User{Name: "Test User"})
	walletID, _ := walletRepo.Save(models.Wallet{Balance: 100.0, UserID: userID})

	tests := []struct {
		name              string
		request           service.CreateDepositRequest
		expectedBalance   float64
		expectedTxnAmount float64
		hasError          bool
		err               error
	}{
		{
			name:              "Valid Deposit",
			request:           service.CreateDepositRequest{WalletID: walletID, Amount: 50.0},
			expectedBalance:   150.0,
			expectedTxnAmount: 50.0,
			hasError:          false,
			err:               nil,
		},
		{
			name:            "Invalid Amount",
			request:         service.CreateDepositRequest{WalletID: walletID, Amount: -10.0},
			expectedBalance: 0,
			hasError:        true,
			err:             errors.ErrInvalidAmount,
		},
		{
			name:            "Non-existent Wallet",
			request:         service.CreateDepositRequest{WalletID: 0, Amount: 50.0},
			expectedBalance: 0,
			hasError:        true,
			err:             errors.ErrResourceNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := walletService.Deposit(tt.request)

			if tt.hasError {
				assert.Error(t, err)
				assert.Equal(t, tt.err, err)

			} else {
				assert.NoError(t, err)

				updatedWallet, err := walletRepo.FindByID(tt.request.WalletID)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBalance, updatedWallet.Balance)
				lastTxn := updatedWallet.Transactions[len(updatedWallet.Transactions)-1]
				assert.Equal(t, "deposit", lastTxn.Type)
				assert.Equal(t, tt.expectedTxnAmount, lastTxn.Amount)
			}
		})
	}
}

func TestWalletService_Withdraw(t *testing.T) {
	_, transactionRepo, walletRepo, userRepo := utils.SetupTestDB()
	walletService := service.NewWalletServiceImpl(transactionRepo, walletRepo)
	userID, _ := userRepo.Save(models.User{Name: "Test User"})
	walletID, _ := walletRepo.Save(models.Wallet{Balance: 100.0, UserID: userID})

	tests := []struct {
		name              string
		request           service.CreateWithdrawalRequest
		expectedBalance   float64
		expectedTxnAmount float64
		hasError          bool
		err               error
	}{
		{
			name:              "Valid Deposit",
			request:           service.CreateWithdrawalRequest{WalletID: walletID, Amount: 50.0},
			expectedBalance:   50.0,
			expectedTxnAmount: 50.0,
			hasError:          false,
			err:               nil,
		},
		{
			name:            "Invalid Amount",
			request:         service.CreateWithdrawalRequest{WalletID: walletID, Amount: -10.0},
			expectedBalance: 0,
			hasError:        true,
			err:             errors.ErrInvalidAmount,
		},
		{
			name:            "Non-existent Wallet",
			request:         service.CreateWithdrawalRequest{WalletID: 0, Amount: 50.0},
			expectedBalance: 0,
			hasError:        true,
			err:             errors.ErrResourceNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := walletService.Withdraw(tt.request)

			if tt.hasError {
				assert.Error(t, err)
				assert.Equal(t, tt.err, err)
			} else {
				assert.NoError(t, err)

				updatedWallet, err := walletRepo.FindByID(tt.request.WalletID)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBalance, updatedWallet.Balance)
				lastTxn := updatedWallet.Transactions[len(updatedWallet.Transactions)-1]
				assert.Equal(t, "withdrawal", lastTxn.Type)
				assert.Equal(t, tt.expectedTxnAmount, lastTxn.Amount)
			}
		})
	}
}

func TestWalletService_Transfer(t *testing.T) {
	_, transactionRepo, walletRepo, userRepo := utils.SetupTestDB()
	walletService := service.NewWalletServiceImpl(transactionRepo, walletRepo)
	userID1, _ := userRepo.Save(models.User{Name: "User 1"})
	walletID1, _ := walletRepo.Save(models.Wallet{Balance: 100.0, UserID: userID1})

	userID2, _ := userRepo.Save(models.User{Name: "User 2"})
	walletID2, _ := walletRepo.Save(models.Wallet{Balance: 100.0, UserID: userID2})

	tests := []struct {
		name           string
		request        service.TransferRequest
		wallet1Balance float64
		wallet2Balance float64
		hasError       bool
		err            error
	}{
		{
			name:           "Valid Transfer",
			request:        service.TransferRequest{FromWalletId: walletID1, ToWalletId: walletID2, Amount: 50.0},
			wallet1Balance: 50.0,
			wallet2Balance: 150.0,
			hasError:       false,
			err:            nil,
		},
		{
			name:           "Invalid Amount",
			request:        service.TransferRequest{FromWalletId: walletID1, ToWalletId: walletID2, Amount: -10.0},
			wallet1Balance: 0,
			wallet2Balance: 0,
			hasError:       true,
			err:            errors.ErrInvalidAmount,
		},
		{
			name:           "Insufficient Funds",
			request:        service.TransferRequest{FromWalletId: walletID1, ToWalletId: walletID2, Amount: 200.0},
			wallet1Balance: 0,
			wallet2Balance: 0,
			hasError:       true,
			err:            errors.ErrInsufficientFunds,
		},
		{
			name:           "Non-existent Wallet",
			request:        service.TransferRequest{FromWalletId: 999, ToWalletId: walletID2, Amount: 30.0},
			wallet1Balance: 0,
			wallet2Balance: 0,
			hasError:       true,
			err:            errors.ErrResourceNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := walletService.Transfer(tt.request)
			if tt.hasError {
				assert.Error(t, err)
				assert.Equal(t, tt.err, err)

			} else {
				assert.NoError(t, err)

				updatedWallet1, err := walletRepo.FindByID(tt.request.FromWalletId)
				assert.NoError(t, err)
				assert.Equal(t, tt.wallet1Balance, updatedWallet1.Balance)

				updatedWallet2, err := walletRepo.FindByID(tt.request.ToWalletId)
				assert.NoError(t, err)
				assert.Equal(t, tt.wallet2Balance, updatedWallet2.Balance)
			}
		})
	}
}
