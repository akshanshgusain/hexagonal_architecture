package service

import "github.com/akshanshgusain/Hexagonal-Architecture/domain"

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
}

// Stub

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s *DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return s.repo.FindAll()
}
