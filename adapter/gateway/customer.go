package gateway

import (
	"context"
	"next-learn-go-gin-sqlc/entity"
	"next-learn-go-gin-sqlc/infrastructure/database/sqlc"
)

type CustomerRepository interface {
	GetAll() ([]*entity.GetAllCustomer, error)
	GetFiltered(query string) ([]*entity.GetFilteredCustomer, error)
	GetAllCount() (int64, error)
}

type customerRepository struct {
	q sqlc.Querier
}

func NewCustomerRepository(q sqlc.Querier) CustomerRepository {
	return &customerRepository{q: q}
}

func (cr *customerRepository) GetAll() ([]*entity.GetAllCustomer, error) {
	cs, err := cr.q.GetAllCustomers(context.Background())
	if err != nil {
		return nil, err
	}
	var customers []*entity.GetAllCustomer
	for _, c := range cs {
		customer, err := entity.ReconstructGetAllCustomer(
			c.ID,
			c.Name,
		)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	return customers, nil
}

func (cr *customerRepository) GetFiltered(query string) ([]*entity.GetFilteredCustomer, error) {
	cs, err := cr.q.GetFilteredCustomers(context.Background(), query)
	if err != nil {
		return nil, err
	}
	var customers []*entity.GetFilteredCustomer
	for _, c := range cs {
		customer, err := entity.ReconstructGetFilteredCustomer(
			c.ID,
			c.Name,
			c.Email,
			c.ImageUrl,
			c.TotalInvoices,
			c.TotalPending,
			c.TotalPaid,
		)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}
	return customers, nil
}

func (cr *customerRepository) GetAllCount() (int64, error) {
	count, err := cr.q.GetCustomerCount(context.Background())
	if err != nil {
		return 0, err
	}
	return count, nil
}
