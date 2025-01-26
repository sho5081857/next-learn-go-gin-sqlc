package handler

import (
	"next-learn-go-gin-sqlc/adapter/controller/presenter"
	"next-learn-go-gin-sqlc/api"
	"next-learn-go-gin-sqlc/pkg/logger"
	"next-learn-go-gin-sqlc/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RevenueHandler struct {
	revenueUseCase usecase.RevenueUseCase
}

func NewRevenueHandler(revenueUseCase usecase.RevenueUseCase) *RevenueHandler {
	return &RevenueHandler{
		revenueUseCase: revenueUseCase,
	}
}

func (r *RevenueHandler) GetAllRevenueList(c *gin.Context) {
	revenues, err := r.revenueUseCase.GetAll()
	if err != nil {
		logger.Error(err.Error())
		c.JSON(presenter.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}
	revenueList := []presenter.AllRevenue{}
	for _, v := range revenues {
		revenueList = append(revenueList, presenter.AllRevenue{
			Month:   v.Month(),
			Revenue: int(v.Revenue()),
		})
	}

	c.JSON(http.StatusOK, presenter.AllRevenueListResponse{
		ApiVersion: api.Version,
		Items:      revenueList,
	})
}
