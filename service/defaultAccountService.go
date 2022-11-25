package service

import (
	"fmt"
	"github.com/akshanshgusain/Hexagonal-Architecture/domain"
	"github.com/akshanshgusain/Hexagonal-Architecture/dto"
	"github.com/akshanshgusain/Hexagonal-Architecture/errs"
	"time"
)

const dbTSLayout = "2006-01-02 15:04:05"

type DefaultAccountService struct {
	repo domain.AccountRepository // a reference to secondary port
}

func (d DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {

	if errr := req.Validate(); errr != nil {
		return nil, errr
	}

	a := domain.Account{
		AccountId:   "",
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02T15:04:05Z07:00"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}

	newAcc, err := d.repo.Save(a)
	if err != nil {
		return nil, err
	}

	response := newAcc.ToNewAccountResponseDto()
	return &response, nil
}

func (d DefaultAccountService) MakeTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {
	if errr := req.Validate(); errr != nil {
		return nil, errr
	}

	// server side validation for checking the available balance in the account

	if req.IsTransactionTypeWithdrawal() {
		account, err := d.repo.FindBy(req.AccountId)
		if err != nil {
			return nil, err
		}
		if !account.CanWithdraw(req.Amount) {
			return nil, errs.NewValidationError("insufficient balance in the account")
		}
	}

	// if all is well, build the domain object & save the transaction
	t := domain.Transaction{
		AccountId:       req.AccountId,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format(dbTSLayout),
	}
	fmt.Println("Before repository saveTransaction call")
	transaction, appError := d.repo.SaveTransaction(t)
	if appError != nil {
		return nil, appError
	}
	fmt.Println("printing transaction")
	fmt.Println(transaction)
	response := transaction.ToDto()
	return &response, nil
}

// Helpers

func NewAccountService(repository domain.AccountRepository) AccountService {
	return &DefaultAccountService{repository}
}
