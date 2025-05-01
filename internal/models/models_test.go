package models_test

// import (
// 	"testing"
// 	"wallet-app/internal/models"

// 	"github.com/stretchr/testify/assert"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// func setupTestDB() (*gorm.DB, error) {
// 	dsn := "host=localhost user=testuser password=testpass dbname=testdb port=5432 sslmode=disable"
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	err = db.AutoMigrate(&models.User{}, &models.Wallet{}, &models.Transaction{})
// 	return db, err
// }

// func TestUserModel(t *testing.T) {
// 	db, err := setupTestDB()
// 	assert.NoError(t, err)

// 	user := models.User{Name: "Test User"}
// 	result := db.Create(&user)
// 	assert.NoError(t, result.Error)
// 	assert.NotZero(t, user.ID)

// 	var fetchedUser models.User
// 	err = db.First(&fetchedUser, user.ID).Error
// 	assert.NoError(t, err)
// 	assert.Equal(t, user.Name, fetchedUser.Name)
// }

// func TestWalletModel(t *testing.T) {
// 	db, err := setupTestDB()
// 	assert.NoError(t, err)

// 	wallet := models.Wallet{UserID: 1, Balance: 100.0}
// 	result := db.Create(&wallet)
// 	assert.NoError(t, result.Error)
// 	assert.NotZero(t, wallet.ID)

// 	var fetchedWallet models.Wallet
// 	err = db.First(&fetchedWallet, wallet.ID).Error
// 	assert.NoError(t, err)
// 	assert.Equal(t, wallet.Balance, fetchedWallet.Balance)
// }

// func TestTransactionModel(t *testing.T) {
// 	db, err := setupTestDB()
// 	assert.NoError(t, err)

// 	wallet := models.Wallet{UserID: 1, Balance: 100.0}
// 	db.Create(&wallet)

// 	transaction := models.Transaction{Amount: 50.0, Type: "credit", WalletID: wallet.ID}
// 	result := db.Create(&transaction)
// 	assert.NoError(t, result.Error)
// 	assert.NotZero(t, transaction.ID)

// 	var fetchedTransaction models.Transaction
// 	err = db.First(&fetchedTransaction, transaction.ID).Error
// 	assert.NoError(t, err)
// 	assert.Equal(t, transaction.Amount, fetchedTransaction.Amount)
// 	assert.Equal(t, transaction.Type, fetchedTransaction.Type)
// }
