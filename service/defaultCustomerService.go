package service

import "github.com/akshanshgusain/Hexagonal-Architecture/domain"

// Stub

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s *DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

func (s *DefaultCustomerService) GetCustomer(id string) (*domain.Customer, error) {
	return s.repo.ById(id)
}

// Helpers

func NewCustomerService(repository domain.CustomerRepository) CustomerService {
	return &DefaultCustomerService{repository}
}
