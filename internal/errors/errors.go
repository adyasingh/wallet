package errors

import (
	"errors"
	"net/http"

	"gorm.io/gorm"
)

var (
	ErrInvalidAmount     = errors.New("amount must be greater than zero")
	ErrInsufficientFunds = errors.New("insufficient funds")
	ErrResourceNotFound  = errors.New("resource not found")
	ErrTechnicalError    = errors.New("technical error")
	ErrInvalidWalletID   = errors.New("invalid wallet id")
	ErrUknown            = errors.New("something went wrong")
)

func HandleDbError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrResourceNotFound
	}
	return ErrTechnicalError
}

func HttpError(err error) (int, string) {
	switch err {
	case ErrInvalidAmount:
		return http.StatusBadRequest, err.Error()
	case ErrInsufficientFunds:
		return http.StatusBadRequest, err.Error()
	case ErrResourceNotFound:
		return http.StatusNotFound, err.Error()
	case ErrTechnicalError:
		return http.StatusInternalServerError, err.Error()
	default:
		return http.StatusInternalServerError, err.Error()
	}
}
