package domain

import "github.com/akshanshgusain/Hexagonal-Architecture/errs"

// Server side PORT/ Secondary PORT

//go:generate mockgen -destination=../mocks/domain/mockAccountRepository.go -package=domain github.com/akshanshgusain/Hexagonal-Architecture/domain AccountRepository
type AccountRepository interface {
	SaveA(account Account) (*Account, *errs.AppError)
	SaveTransaction(transaction Transaction) (*Transaction, *errs.AppError)
	FindBy(accountId string) (*Account, *errs.AppError)
}
