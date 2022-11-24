package dto

import (
	"github.com/akshanshgusain/Hexagonal-Architecture/errs"
	"strings"
)

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (n *NewAccountRequest) Validate() *errs.AppError {
	if n.Amount < 5000 {
		return errs.NewValidationError("amount is below the minimum amount")
	}
	if strings.ToLower(n.AccountType) != "saving" && strings.ToLower(n.AccountType) != "checking" {
		return errs.NewValidationError("account type should be either saving or checking")
	}
	return nil
}
