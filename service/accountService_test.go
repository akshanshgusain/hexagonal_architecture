package service

import (
	"github.com/akshanshgusain/Hexagonal-Architecture/dto"
	"testing"
)

func Test_should_return_a_validation_error_response_when_the_request_is_not_validated(t *testing.T) {
	// Arrange
	request := dto.NewAccountRequest{
		CustomerId:  "100",
		AccountType: "saving",
		Amount:      0,
	}
	service := NewAccountService(nil)
	// Act
	_, appError := service.NewAccount(request)
	// Assert
	if appError == nil {
		t.Error("failed while testing the new account validation")
	}
}
