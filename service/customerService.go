package service

import "github.com/akshanshgusain/Hexagonal-Architecture/domain"

// client side PORT/ Primary PORT

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
	GetCustomer(id string) (*domain.Customer, error)
}
