package service

import (
	"github.com/akshanshgusain/Hexagonal-Architecture/domain"
	"github.com/akshanshgusain/Hexagonal-Architecture/dto"
	"github.com/akshanshgusain/Hexagonal-Architecture/errs"
)

// Stub

//go:generate mockgen -destination=../mocks/service/mockCustomerService.go -package=service github.com/akshanshgusain/Hexagonal-Architecture/service CustomerService
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s *DefaultCustomerService) GetAllCustomers() ([]dto.CustomerResponse, *errs.AppError) {
	cs, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	var response = make([]dto.CustomerResponse, 0)
	for _, cr := range cs {
		response = append(response, cr.ToDto())
	}
	return response, nil
}

func (s *DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}
	response := c.ToDto()
	return &response, nil
}

// Helpers

func NewCustomerService(repository domain.CustomerRepository) CustomerService {
	return &DefaultCustomerService{repository}
}
