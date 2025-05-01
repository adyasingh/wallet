package repository

import (
	"wallet-app/internal/errors"
	"wallet-app/internal/models"

	"gorm.io/gorm"
)

type WalletRepository interface {
	Save(models.Wallet) (uint, error)
	FindByID(id uint) (models.Wallet, error)
	UpdateBalance(id uint, amount float64) error
}

type WalletRepositoryImpl struct {
	Db *gorm.DB
}

func NewWalletRepositoryImpl(Db *gorm.DB) WalletRepositoryImpl {
	return WalletRepositoryImpl{Db: Db}
}

func (u WalletRepositoryImpl) Save(wallet models.Wallet) (uint, error) {
	result := u.Db.Create(&wallet)
	if err := result.Error; err != nil {
		return 0, errors.HandleDbError(err)
	}
	return wallet.ID, nil
}

func (w WalletRepositoryImpl) FindByID(id uint) (models.Wallet, error) {
	var wallet models.Wallet
	if err := w.Db.Preload("Transactions").First(&wallet, id).Error; err != nil {
		return wallet, errors.HandleDbError(err)
	}
	return wallet, nil
}

func (w WalletRepositoryImpl) UpdateBalance(id uint, amount float64) error {
	var wallet models.Wallet
	if err := w.Db.First(&wallet, id).Error; err != nil {
		return errors.HandleDbError(err)
	}
	wallet.Balance = amount
	if err := w.Db.Save(&wallet).Error; err != nil {
		return errors.HandleDbError(err)
	}
	return nil
}
