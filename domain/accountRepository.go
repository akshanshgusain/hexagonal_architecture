package domain

import "github.com/akshanshgusain/Hexagonal-Architecture/errs"

type Account struct {
	AccountId   string
	CustomerId  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

// Server side PORT/ Secondary PORT

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}
