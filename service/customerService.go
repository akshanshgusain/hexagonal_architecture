package service

import "github.com/akshanshgusain/Hexagonal-Architecture/domain"

// client side PORT

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
}
