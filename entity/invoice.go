package entity

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"

)

const (
	Paid   InvoiceStatus = "paid"
	Unpaid InvoiceStatus = "pending"
)

type PaidCount int64
type PendingCount int64

type InvoiceStatus string

type Invoice struct {
	id         uuid.UUID
	amount     int32
	status     string
	date       pgtype.Date
	customerID uuid.UUID
}

func (i *Invoice) ID() uuid.UUID {
	return i.id
}

func (i *Invoice) Amount() int32 {
	return i.amount
}

func (i *Invoice) Status() string {
	return i.status
}

func (i *Invoice) Date() pgtype.Date {
	return i.date
}

func (i *Invoice) CustomerID() uuid.UUID {
	return i.customerID
}

func (g *GetLatestInvoice) ID() uuid.UUID {
	return g.id
}

func (g *GetLatestInvoice) Name() string {
	return g.name
}

func (g *GetLatestInvoice) Email() string {
	return g.email
}

func (g *GetLatestInvoice) ImageUrl() string {
	return g.imageUrl
}

func (g *GetLatestInvoice) Amount() int32 {
	return g.amount
}

func (g *GetFilteredInvoice) ID() uuid.UUID {
	return g.id
}

func (g *GetFilteredInvoice) Name() string {
	return g.name
}

func (g *GetFilteredInvoice) Email() string {
	return g.email
}

func (g *GetFilteredInvoice) ImageUrl() string {
	return g.imageUrl
}

func (g *GetFilteredInvoice) Amount() int32 {
	return g.amount
}

func (g *GetFilteredInvoice) Date() pgtype.Date {
	return g.date
}

func (g *GetFilteredInvoice) Status() string {
	return g.status
}

func (g *GetInvoiceById) ID() uuid.UUID {
	return g.id
}

func (g *GetInvoiceById) CustomerID() uuid.UUID {
	return g.customerID
}

func (g *GetInvoiceById) Amount() int32 {
	return g.amount
}

func (g *GetInvoiceById) Status() string {
	return g.status
}


type GetLatestInvoice struct {
	id       uuid.UUID
	name     string
	imageUrl string
	email    string
	amount   int32
}

type GetFilteredInvoice struct {
	id       uuid.UUID
	name     string
	email    string
	imageUrl string
	amount   int32
	date     pgtype.Date
	status   string
}

type GetInvoiceById struct {
	id         uuid.UUID
	customerID uuid.UUID
	amount     int32
	status     string
}

func (is *InvoiceStatus) IsValid() bool {
	return *is == Paid || *is == Unpaid
}

func ReconstructInvoice(
	id uuid.UUID,
	amount int32,
	status InvoiceStatus,
	date pgtype.Date,
	customerID uuid.UUID,
) (*Invoice, error) {
	return newInvoice(
		id,
		amount,
		status,
		date,
		customerID,
	)
}

func NewInvoice(
	id uuid.UUID,
	amount int32,
	status InvoiceStatus,
	date pgtype.Date,
	customerID uuid.UUID,
) (*Invoice, error) {
	return newInvoice(
		id,
		amount,
		status,
		date,
		customerID,
	)
}

func newInvoice(
	id uuid.UUID,
	amount int32,
	status InvoiceStatus,
	date pgtype.Date,
	customerID uuid.UUID,
) (*Invoice, error) {
	if !status.IsValid() {
		return nil, errors.New("invalid status")
	}
	return &Invoice{
		id:         id,
		amount:     amount,
		status:     string(status),
		date:       date,
		customerID: customerID,
	}, nil
}

func ReconstructGetLatestInvoice(
	id uuid.UUID,
	name string,
	imageUrl string,
	email string,
	amount int32,
) (*GetLatestInvoice, error) {
	return newGetLatestInvoice(
		id,
		name,
		imageUrl,
		email,
		amount,
	)
}

func newGetLatestInvoice(
	id uuid.UUID,
	name string,
	imageUrl string,
	email string,
	amount int32,
) (*GetLatestInvoice, error) {
	return &GetLatestInvoice{
		id:       id,
		name:     name,
		imageUrl: imageUrl,
		email:    email,
		amount:   amount,
	}, nil
}

func ReconstructGetFilteredInvoice(
	id uuid.UUID,
	name string,
	email string,
	imageUrl string,
	amount int32,
	date pgtype.Date,
	status InvoiceStatus,
) (*GetFilteredInvoice, error) {
	return newGetFilteredInvoice(
		id,
		name,
		email,
		imageUrl,
		amount,
		date,
		status,
	)
}

func newGetFilteredInvoice(
	id uuid.UUID,
	name string,
	email string,
	imageUrl string,
	amount int32,
	date pgtype.Date,
	status InvoiceStatus,
) (*GetFilteredInvoice, error) {
	if !status.IsValid() {
		return nil, errors.New("invalid status")
	}
	return &GetFilteredInvoice{
		id:       id,
		name:     name,
		email:    email,
		imageUrl: imageUrl,
		amount:   amount,
		date:     date,
		status:   string(status),
	}, nil
}

func ReconstructGetInvoiceById(
	id uuid.UUID,
	customerID uuid.UUID,
	amount int32,
	status InvoiceStatus,
) (*GetInvoiceById, error) {
	return newGetInvoiceById(
		id,
		customerID,
		amount,
		status,
	)
}

func newGetInvoiceById(
	id uuid.UUID,
	customerID uuid.UUID,
	amount int32,
	status InvoiceStatus,
) (*GetInvoiceById, error) {
	if !status.IsValid() {
		return nil, errors.New("invalid status")
	}
	return &GetInvoiceById{
		id:         id,
		customerID: customerID,
		amount:     amount,
		status:     string(status),
	}, nil
}
