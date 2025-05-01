package errors_test

import (
	"errors"
	"net/http"
	"testing"

	errInternal "wallet-app/internal/errors"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestHandleDbError(t *testing.T) {
	tests := []struct {
		inputError    error
		expectedError error
	}{
		{gorm.ErrRecordNotFound, errInternal.ErrResourceNotFound},
		{errors.New("some other error"), errInternal.ErrTechnicalError},
	}

	for _, test := range tests {
		result := errInternal.HandleDbError(test.inputError)
		assert.Equal(t, test.expectedError, result)
	}
}

func TestHttpError(t *testing.T) {
	tests := []struct {
		err          error
		expectedCode int
		expectedMsg  string
	}{
		{errInternal.ErrInvalidAmount, http.StatusBadRequest, errInternal.ErrInvalidAmount.Error()},
		{errInternal.ErrInsufficientFunds, http.StatusBadRequest, errInternal.ErrInsufficientFunds.Error()},
		{errInternal.ErrResourceNotFound, http.StatusNotFound, errInternal.ErrResourceNotFound.Error()},
		{errInternal.ErrTechnicalError, http.StatusInternalServerError, errInternal.ErrTechnicalError.Error()},
		{errInternal.ErrUknown, http.StatusInternalServerError, errInternal.ErrUknown.Error()},
	}

	for _, test := range tests {
		code, msg := errInternal.HttpError(test.err)
		assert.Equal(t, test.expectedCode, code)
		assert.Equal(t, test.expectedMsg, msg)
	}
}
