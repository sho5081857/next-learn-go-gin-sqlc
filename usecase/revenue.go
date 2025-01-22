package usecase

import (
	"next-learn-go-gin-sqlc/adapter/gateway"
	"next-learn-go-gin-sqlc/entity"
)

type RevenueUseCase interface {
	GetAll() ([]*entity.Revenue, error)
}

type revenueUseCase struct {
	revenueRepository gateway.RevenueRepository
}

func NewRevenueUseCase(revenueRepository gateway.RevenueRepository) RevenueUseCase {
	return &revenueUseCase{
		revenueRepository: revenueRepository,
	}
}

func (r *revenueUseCase) GetAll() ([]*entity.Revenue, error) {
	return r.revenueRepository.GetAll()
}
