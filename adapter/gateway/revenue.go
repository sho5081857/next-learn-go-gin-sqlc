package gateway

import (
	"context"
	"next-learn-go-gin-sqlc/entity"
	"next-learn-go-gin-sqlc/infrastructure/database/sqlc"
)

type RevenueRepository interface {
	GetAll() ([]*entity.Revenue, error)
}

type revenueRepository struct {
	q sqlc.Querier
}

func NewRevenueRepository(q sqlc.Querier) RevenueRepository {
	return &revenueRepository{q: q}
}

func (rr *revenueRepository) GetAll() ([]*entity.Revenue, error) {
	rs, err := rr.q.GetAllRevenue(context.Background())
	if err != nil {
		return nil, err
	}
	var revenues []*entity.Revenue
	for _, r := range rs {
		revenue, err := entity.ReconstructRevenue(
			r.Month,
			r.Revenue,
		)
		if err != nil {
			return nil, err
		}
		revenues = append(revenues, revenue)
	}
	return revenues, nil
}
