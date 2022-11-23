package service

import (
	"github.com/akshanshgusain/Hexagonal-Architecture/domain"
	"github.com/akshanshgusain/Hexagonal-Architecture/errs"
)

// client side PORT/ Primary PORT

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
	GetCustomer(id string) (*domain.Customer, *errs.AppError)
}
