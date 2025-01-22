package usecase

import (
	"next-learn-go-gin-sqlc/adapter/gateway"
	"next-learn-go-gin-sqlc/entity"
)

type CustomerUseCase interface {
	GetAll() ([]*entity.GetAllCustomer, error)
	GetFiltered(query string) ([]*entity.GetFilteredCustomer, error)
	GetAllCount() (int64, error)
}

type customerUseCase struct {
	customerRepository gateway.CustomerRepository
}

func NewCustomerUseCase(customerRepository gateway.CustomerRepository) CustomerUseCase {
	return &customerUseCase{
		customerRepository: customerRepository,
	}
}

func (c *customerUseCase) GetAll() ([]*entity.GetAllCustomer, error) {
	return c.customerRepository.GetAll()
}

func (c *customerUseCase) GetFiltered(query string) ([]*entity.GetFilteredCustomer, error) {
	return c.customerRepository.GetFiltered(query)
}

func (c *customerUseCase) GetAllCount() (int64, error) {
	return c.customerRepository.GetAllCount()
}
