package repository_test

import (
	"testing"
	"wallet-app/internal/models"
	"wallet-app/internal/utils" // Updated import

	"github.com/stretchr/testify/assert"
)

func TestWalletRepository_Save(t *testing.T) {
	db, _, walletRepo, userRepo := utils.SetupTestDB()
	user := models.User{Name: "Test User"}
	id, _ := userRepo.Save(user)
	wallet := models.Wallet{Balance: 100.0, UserID: id}
	walletID, err := walletRepo.Save(wallet)

	assert.NoError(t, err)
	var savedWallet models.Wallet
	err = db.First(&savedWallet, walletID).Error
	assert.NoError(t, err)
	assert.Equal(t, wallet.Balance, savedWallet.Balance)
	assert.Equal(t, wallet.UserID, savedWallet.UserID)
}

func TestWalletRepository_FindByID(t *testing.T) {
	_, _, walletRepo, userRepo := utils.SetupTestDB()
	user := models.User{Name: "Test User"}
	id, _ := userRepo.Save(user)
	wallet := models.Wallet{Balance: 100.0, UserID: id}
	walletID, _ := walletRepo.Save(wallet)

	foundWallet, err := walletRepo.FindByID(walletID)

	assert.NoError(t, err)
	assert.Equal(t, wallet.Balance, foundWallet.Balance)
	assert.Equal(t, wallet.UserID, foundWallet.UserID)
}

func TestWalletRepository_UpdateBalance(t *testing.T) {
	_, _, walletRepo, userRepo := utils.SetupTestDB()
	user := models.User{Name: "Test User"}
	id, _ := userRepo.Save(user)
	wallet := models.Wallet{Balance: 100.0, UserID: id}
	walletID, _ := walletRepo.Save(wallet)

	newBalance := 300.0
	err := walletRepo.UpdateBalance(walletID, newBalance)

	assert.NoError(t, err)

	updatedWallet, err := walletRepo.FindByID(walletID)
	assert.NoError(t, err)
	assert.Equal(t, newBalance, updatedWallet.Balance)
}
