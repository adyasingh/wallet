package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name   string
	Wallet Wallet `gorm:"foreignKey:UserID"`
}

type Wallet struct {
	gorm.Model
	UserID       uint
	Balance      float64
	Transactions []Transaction `gorm:"foreignKey:WalletID"`
	Transfers    []Transaction `gorm:"foreignKey:TransferWalletID"`
}

type Transaction struct {
	gorm.Model
	Amount           float64
	Type             string
	WalletID         uint
	TransferWalletID *uint
}
