package handlers

import "wallet-app/internal/service"

type WalletHandler struct {
	walletService service.WalletService
}

func NewWalletHandler(service service.WalletService) *WalletHandler {
	return &WalletHandler{walletService: service}
}
