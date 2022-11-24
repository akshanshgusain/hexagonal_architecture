package service

import (
	"github.com/akshanshgusain/Hexagonal-Architecture/domain"
	"github.com/akshanshgusain/Hexagonal-Architecture/dto"
	"github.com/akshanshgusain/Hexagonal-Architecture/errs"
	"time"
)

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

// Helpers

func NewAccountService(repository domain.AccountRepository) AccountService {
	return &DefaultAccountService{repository}
}
