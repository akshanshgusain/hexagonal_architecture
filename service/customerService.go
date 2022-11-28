package service

import (
	"github.com/akshanshgusain/Hexagonal-Architecture/dto"
	"github.com/akshanshgusain/Hexagonal-Architecture/errs"
)

// client side PORT/ Primary PORT

type CustomerService interface {
	GetAllCustomers() ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError)
}
