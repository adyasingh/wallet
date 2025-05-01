package repository_test

import (
	"testing"
	"wallet-app/internal/models"
	"wallet-app/internal/utils"

	"github.com/stretchr/testify/assert"
)

func TestTransactionRepository_Save(t *testing.T) {
	db, transactionRepo, walletRepo, userRepo := utils.SetupTestDB()

	userID, _ := userRepo.Save(models.User{Name: "Test User"})
	walletID, _ := walletRepo.Save(models.Wallet{Balance: 100.0, UserID: userID})

	transaction := models.Transaction{Amount: 50.0, WalletID: walletID}

	txnID, err := transactionRepo.Save(transaction)
	assert.NoError(t, err)

	var savedTransaction models.Transaction
	err = db.First(&savedTransaction, txnID).Error
	assert.NoError(t, err)
	assert.Equal(t, transaction.Amount, savedTransaction.Amount)
	assert.Equal(t, transaction.WalletID, savedTransaction.WalletID)
}
