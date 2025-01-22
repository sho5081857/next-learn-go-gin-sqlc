package gateway

import (
	"context"
	"next-learn-go-gin-sqlc/entity"
	"next-learn-go-gin-sqlc/infrastructure/database/sqlc"

	"github.com/google/uuid"
)

type InvoiceRepository interface {
	GetLatest() ([]*entity.GetLatestInvoice, error)
	GetFiltered(query string, offset, limit int32) ([]*entity.GetFilteredInvoice, int64, error)
	GetAllCount() (int64, error)
	GetStatusCount() (entity.PaidCount, entity.PendingCount, error)
	GetById(ID uuid.UUID) (*entity.GetInvoiceById, error)
	Create(invoice *entity.Invoice) (*entity.Invoice, error)
	Update(invoice *entity.Invoice, ID uuid.UUID) (*entity.Invoice, error)
	Delete(ID uuid.UUID) error
}

type invoiceRepository struct {
	q sqlc.Querier
}

func NewInvoiceRepository(q sqlc.Querier) InvoiceRepository {
	return &invoiceRepository{q: q}
}

func (ir *invoiceRepository) GetLatest() ([]*entity.GetLatestInvoice, error) {
	is, err := ir.q.GetLatestInvoices(context.Background())
	if err != nil {
		return nil, err
	}
	var invoices []*entity.GetLatestInvoice
	for _, i := range is {
		invoice, err := entity.ReconstructGetLatestInvoice(
			i.ID,
			i.Name,
			i.ImageUrl,
			i.Email,
			i.Amount,
		)
		if err != nil {
			return nil, err
		}
		invoices = append(invoices, invoice)
	}
	return invoices, nil
}

func (ir *invoiceRepository) GetFiltered(query string, offset, limit int32) ([]*entity.GetFilteredInvoice, int64, error) {
	is, err := ir.q.GetFilteredInvoices(context.Background(), sqlc.GetFilteredInvoicesParams{Name: query, Offset: offset, Limit: limit})
	if err != nil {
		return nil, 0, err
	}
	var invoices []*entity.GetFilteredInvoice
	for _, i := range is {
		invoice, err := entity.ReconstructGetFilteredInvoice(
			i.ID,
			i.Name,
			i.Email,
			i.ImageUrl,
			i.Amount,
			i.Date,
			entity.InvoiceStatus(i.Status),
		)
		if err != nil {
			return nil, 0, err
		}
		invoices = append(invoices, invoice)
	}

	totalCount, err := ir.q.GetInvoicesPages(context.Background(), query)
	if err != nil {
		return nil, 0, err
	}
	return invoices, totalCount, nil
}

func (ir *invoiceRepository) GetAllCount() (int64, error) {
	count, err := ir.q.GetInvoiceCount(context.Background())
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (ir *invoiceRepository) GetStatusCount() (entity.PaidCount, entity.PendingCount, error) {
	var paid entity.PaidCount
	var pending entity.PendingCount
	counts, err := ir.q.GetInvoiceStatusCount(context.Background())
	if err != nil {
		return paid, pending, err
	}
	paid = entity.PaidCount(counts.Paid)
	pending = entity.PendingCount(counts.Pending)

	return paid, pending, nil
}

func (ir *invoiceRepository) GetById(ID uuid.UUID) (*entity.GetInvoiceById, error) {
	i, err := ir.q.GetInvoiceById(context.Background(), ID)
	if err != nil {
		return nil, err
	}
	invoice, err := entity.ReconstructGetInvoiceById(
		i.ID,
		i.CustomerID,
		i.Amount,
		entity.InvoiceStatus(i.Status),
	)
	if err != nil {
		return nil, err
	}

	return invoice, nil
}

func (ir *invoiceRepository) Create(invoice *entity.Invoice) (*entity.Invoice, error) {
	i, err := ir.q.CreateInvoice(context.Background(), sqlc.CreateInvoiceParams{
		Amount:     invoice.Amount(),
		Status:     invoice.Status(),
		Date:       invoice.Date(),
		CustomerID: invoice.CustomerID(),
	})
	if err != nil {
		return nil, err
	}
	createdInvoice, err := entity.ReconstructInvoice(
		i.ID,
		i.Amount,
		entity.InvoiceStatus(i.Status),
		i.Date,
		i.CustomerID,
	)
	if err != nil {
		return nil, err
	}
	return createdInvoice, nil
}

func (ir *invoiceRepository) Update(invoice *entity.Invoice, ID uuid.UUID) (*entity.Invoice, error) {

	i, err := ir.q.UpdateInvoice(context.Background(), sqlc.UpdateInvoiceParams{
		ID:         ID,
		Amount:     invoice.Amount(),
		Status:     invoice.Status(),
		CustomerID: invoice.CustomerID(),
	})
	if err != nil {
		return nil, err
	}
	updatedInvoice, err := entity.ReconstructInvoice(
		i.ID,
		i.Amount,
		entity.InvoiceStatus(i.Status),
		i.Date,
		i.CustomerID,
	)

	if err != nil {
		return nil, err
	}
	return updatedInvoice, nil
}

func (ir *invoiceRepository) Delete(ID uuid.UUID) error {
	if err := ir.q.DeleteInvoice(context.Background(), ID); err != nil {
		return err
	}
	return nil
}
