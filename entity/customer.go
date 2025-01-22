package entity

import "github.com/google/uuid"


type Customer struct {
	id       uuid.UUID
	name     string
	email    string
	imageUrl string
}

type GetAllCustomer struct {
	id   uuid.UUID
	name string
}

type GetFilteredCustomer struct {
	id            uuid.UUID
	name          string
	email         string
	imageUrl      string
	totalInvoices int64
	totalPending  int64
	totalPaid     int64
}

func (c *Customer) ID() uuid.UUID {
	return c.id
}

func (c *Customer) Name() string {
	return c.name
}

func (c *Customer) Email() string {
	return c.email
}

func (c *Customer) ImageUrl() string {
	return c.imageUrl
}

func (g *GetAllCustomer) ID() uuid.UUID {
	return g.id
}

func (g *GetAllCustomer) Name() string {
	return g.name
}

func (g *GetFilteredCustomer) ID() uuid.UUID {
	return g.id
}

func (g *GetFilteredCustomer) Name() string {
	return g.name
}

func (g *GetFilteredCustomer) Email() string {
	return g.email
}

func (g *GetFilteredCustomer) ImageUrl() string {
	return g.imageUrl
}

func (g *GetFilteredCustomer) TotalInvoices() int64 {
	return g.totalInvoices
}

func (g *GetFilteredCustomer) TotalPending() int64 {
	return g.totalPending
}

func (g *GetFilteredCustomer) TotalPaid() int64 {
	return g.totalPaid
}

func ReconstructCustomer(
	id uuid.UUID,
	name string,
	email string,
	imageUrl string,
) (*Customer, error) {
	return newCustomer(
		id,
		name,
		email,
		imageUrl,
	)
}

func newCustomer(
	id uuid.UUID,
	name string,
	email string,
	imageUrl string,
) (*Customer, error) {
	return &Customer{
		id:       id,
		name:     name,
		email:    email,
		imageUrl: imageUrl,
	}, nil
}

func ReconstructGetAllCustomer(
	id uuid.UUID,
	name string,
) (*GetAllCustomer, error) {
	return newGetAllCustomer(
		id,
		name,
	)
}

func newGetAllCustomer(
	id uuid.UUID,
	name string,
) (*GetAllCustomer, error) {
	return &GetAllCustomer{
		id:   id,
		name: name,
	}, nil
}

func ReconstructGetFilteredCustomer(
	id uuid.UUID,
	name string,
	email string,
	imageUrl string,
	totalInvoices int64,
	totalPending int64,
	totalPaid int64,
) (*GetFilteredCustomer, error) {
	return newGetFilteredCustomer(
		id,
		name,
		email,
		imageUrl,
		totalInvoices,
		totalPending,
		totalPaid,
	)
}

func newGetFilteredCustomer(
	id uuid.UUID,
	name string,
	email string,
	imageUrl string,
	totalInvoices int64,
	totalPending int64,
	totalPaid int64,
) (*GetFilteredCustomer, error) {
	return &GetFilteredCustomer{
		id:            id,
		name:          name,
		email:         email,
		imageUrl:      imageUrl,
		totalInvoices: totalInvoices,
		totalPending:  totalPending,
		totalPaid:     totalPaid,
	}, nil
}
