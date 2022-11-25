package domain

import (
	"github.com/akshanshgusain/Hexagonal-Architecture/errs"
)

// Server side PORT/ Secondary PORT

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
	SaveTransaction(transaction Transaction) (*Transaction, *errs.AppError)
	FindBy(accountId string) (*Account, *errs.AppError)
}
