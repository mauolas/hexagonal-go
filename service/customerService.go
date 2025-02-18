package service

import (
	"example.com/hexagonal/domain"
	"example.com/hexagonal/dto"
	"example.com/hexagonal/errs"
)

type CustomerService interface {
	GetAllCustomers() ([]dto.CustomerResponse, error)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
	GetCustomerByStatus(string) ([]domain.Customer, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers() ([]dto.CustomerResponse, error) {
	c, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	var response []dto.CustomerResponse
	for _, v := range c {
		response = append(response, v.ToDto())
	}
	return response, nil
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}
	response := c.ToDto()
	return &response, nil
}

func (s DefaultCustomerService) GetCustomerByStatus(status string) ([]domain.Customer, *errs.AppError) {
	return s.repo.FindByStatus(status)
}

func NewCustomerService(repo domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repo}
}
