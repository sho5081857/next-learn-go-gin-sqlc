package usecase

import (
	"next-learn-go-gin-sqlc/adapter/gateway"
	"next-learn-go-gin-sqlc/entity"

	"github.com/google/uuid"

)

type InvoiceUseCase interface {
	GetLatest() ([]*entity.GetLatestInvoice, error)
	GetFiltered(query string, offset, limit int32) ([]*entity.GetFilteredInvoice, int64, error)
	GetAllCount() (int64, error)
	GetStatusCount() (entity.PaidCount, entity.PendingCount, error)
	GetById(ID uuid.UUID) (*entity.GetInvoiceById, error)
	Create(invoice *entity.Invoice) (*entity.Invoice, error)
	Update(invoice *entity.Invoice, ID uuid.UUID) (*entity.Invoice, error)
	Delete(ID uuid.UUID) error
}

type invoiceUseCase struct {
	invoiceRepository gateway.InvoiceRepository
}

func NewInvoiceUseCase(invoiceRepository gateway.InvoiceRepository) InvoiceUseCase {
	return &invoiceUseCase{
		invoiceRepository: invoiceRepository,
	}
}

func (i *invoiceUseCase) GetLatest() ([]*entity.GetLatestInvoice, error) {
	return i.invoiceRepository.GetLatest()
}

func (i *invoiceUseCase) GetFiltered(query string, offset, limit int32) ([]*entity.GetFilteredInvoice, int64, error) {
	return i.invoiceRepository.GetFiltered(query, offset, limit)
}

func (i *invoiceUseCase) GetAllCount() (int64, error) {
	return i.invoiceRepository.GetAllCount()
}

func (i *invoiceUseCase) GetStatusCount() (entity.PaidCount, entity.PendingCount, error) {
	return i.invoiceRepository.GetStatusCount()
}

func (i *invoiceUseCase) GetById(ID uuid.UUID) (*entity.GetInvoiceById, error) {
	return i.invoiceRepository.GetById(ID)
}

func (i *invoiceUseCase) Create(invoice *entity.Invoice) (*entity.Invoice, error) {
	return i.invoiceRepository.Create(invoice)
}

func (i *invoiceUseCase) Update(invoice *entity.Invoice, ID uuid.UUID) (*entity.Invoice, error) {
	return i.invoiceRepository.Update(invoice, ID)
}

func (i *invoiceUseCase) Delete(ID uuid.UUID) error {
	return i.invoiceRepository.Delete(ID)
}
