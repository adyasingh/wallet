package service

import (
	"wallet-app/internal/errors"
	"wallet-app/internal/models"
	"wallet-app/internal/repository"
)

type WalletService interface {
	Deposit(CreateDepositRequest) error
	GetBalance(id uint) (float64, error)
	GetTransactions(id uint) ([]models.Transaction, error)
	Withdraw(CreateWithdrawalRequest) error
	Transfer(TransferRequest) error
}

type WalletServiceImpl struct {
	TransactionRepository repository.TransactionRepository
	WalletRepository      repository.WalletRepository
}

func NewWalletServiceImpl(transactionRespository repository.TransactionRepository, walletRepository repository.WalletRepository) WalletServiceImpl {
	return WalletServiceImpl{
		TransactionRepository: transactionRespository,
		WalletRepository:      walletRepository,
	}
}

type CreateDepositRequest struct {
	WalletID uint
	Amount   float64 `json:"amount" binding:"required"`
}

func (w WalletServiceImpl) Deposit(deposit CreateDepositRequest) error {
	if deposit.Amount <= 0 {
		return errors.ErrInvalidAmount
	}

	wallet, err := w.WalletRepository.FindByID(deposit.WalletID)
	if err != nil {
		return err
	}

	transaction := models.Transaction{
		WalletID: uint(deposit.WalletID),
		Amount:   deposit.Amount,
		Type:     "deposit",
	}

	_, err = w.TransactionRepository.Save(transaction)
	if err != nil {
		return err
	}

	remainingBalance := wallet.Balance + deposit.Amount
	return w.WalletRepository.UpdateBalance(deposit.WalletID, remainingBalance)
}

func (w WalletServiceImpl) GetBalance(id uint) (float64, error) {
	wallet, err := w.WalletRepository.FindByID(id)
	if err != nil {
		return 0, err
	}
	return wallet.Balance, nil
}

func (w WalletServiceImpl) GetTransactions(id uint) ([]models.Transaction, error) {
	wallet, err := w.WalletRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return wallet.Transactions, nil
}

type CreateWithdrawalRequest struct {
	WalletID uint
	Amount   float64 `json:"amount" binding:"required"`
}

func (w WalletServiceImpl) Withdraw(withdrawal CreateWithdrawalRequest) error {
	if withdrawal.Amount <= 0 {
		return errors.ErrInvalidAmount
	}

	wallet, err := w.WalletRepository.FindByID(withdrawal.WalletID)
	if err != nil {
		return err
	}

	remainingBalance := wallet.Balance - withdrawal.Amount
	if remainingBalance < 0 {
		return errors.ErrInsufficientFunds
	}

	transaction := models.Transaction{
		WalletID: uint(withdrawal.WalletID),
		Amount:   withdrawal.Amount,
		Type:     "withdrawal",
	}

	_, err = w.TransactionRepository.Save(transaction)
	if err != nil {
		return err
	}

	return w.WalletRepository.UpdateBalance(withdrawal.WalletID, remainingBalance)
}

type TransferRequest struct {
	FromWalletId uint
	ToWalletId   uint    `json:"to_wallet_id" binding:"required"`
	Amount       float64 `json:"amount" binding:"required"`
}

func (w WalletServiceImpl) Transfer(transfer TransferRequest) error {
	if transfer.Amount <= 0 {
		return errors.ErrInvalidAmount
	}

	fromWallet, err := w.WalletRepository.FindByID(transfer.FromWalletId)
	if err != nil {
		return err
	}
	remainingBalance := fromWallet.Balance - transfer.Amount
	if remainingBalance < 0 {
		return errors.ErrInsufficientFunds
	}

	toWallet, err := w.WalletRepository.FindByID(transfer.ToWalletId)
	if err != nil {
		return err
	}

	outgoingTxn := models.Transaction{
		WalletID:         fromWallet.ID,
		Amount:           transfer.Amount,
		Type:             "transfer_outgoing",
		TransferWalletID: &toWallet.ID,
	}
	_, err = w.TransactionRepository.Save(outgoingTxn)
	if err != nil {
		return err
	}
	err = w.WalletRepository.UpdateBalance(transfer.FromWalletId, remainingBalance)
	if err != nil {
		return err
	}

	incomingTxn := models.Transaction{
		WalletID:         toWallet.ID,
		Amount:           transfer.Amount,
		Type:             "transfer_incoming",
		TransferWalletID: &fromWallet.ID,
	}
	_, err = w.TransactionRepository.Save(incomingTxn)
	if err != nil {
		return err
	}

	return w.WalletRepository.UpdateBalance(toWallet.ID, toWallet.Balance+transfer.Amount)
}
