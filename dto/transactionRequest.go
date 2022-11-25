package dto

import (
	"github.com/akshanshgusain/Hexagonal-Architecture/errs"
)

const WITHDRAWAL = "withdrawal"
const DEPOSIT = "deposit"

type TransactionRequest struct {
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
	CustomerId      string  `json:"-"`
}

func (r *TransactionRequest) IsTransactionTypeWithdrawal() bool {
	return r.TransactionType == WITHDRAWAL
}

func (r *TransactionRequest) IsTransactionTypeDeposit() bool {
	return r.TransactionType == DEPOSIT
}

func (t *TransactionRequest) Validate() *errs.AppError {
	if t.TransactionType != WITHDRAWAL && t.TransactionType != DEPOSIT {
		return errs.NewValidationError("transaction type can only be deposit or withdrawal")
	}
	if t.Amount < 0 {
		return errs.NewValidationError("amount cannot be less than zero")
	}

	return nil
}
