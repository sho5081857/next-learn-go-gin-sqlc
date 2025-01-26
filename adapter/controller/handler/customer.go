package handler

import (
	"net/http"
	"next-learn-go-gin-sqlc/adapter/controller/presenter"
	"next-learn-go-gin-sqlc/api"
	"next-learn-go-gin-sqlc/pkg/logger"
	"next-learn-go-gin-sqlc/usecase"

	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	customerUseCase usecase.CustomerUseCase
}

func NewCustomerHandler(customerUseCase usecase.CustomerUseCase) *CustomerHandler {
	return &CustomerHandler{
		customerUseCase: customerUseCase,
	}
}

func (ch *CustomerHandler) GetAllCustomersList(c *gin.Context) {
	customers, err := ch.customerUseCase.GetAll()
	if err != nil {
		logger.Error(err.Error())
		c.JSON(presenter.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}
	customerList := []presenter.AllCustomer{}
	for _, v := range customers {
		customerList = append(customerList, presenter.AllCustomer{
			Id:   v.ID(),
			Name: v.Name(),
		})
	}

	c.JSON(http.StatusOK, presenter.AllCustomersListResponse{
		ApiVersion: api.Version,
		Items:      customerList,
	})
}

func (ch *CustomerHandler) GetFilteredCustomersList(c *gin.Context, params presenter.GetFilteredCustomersListParams) {
	customers, err := ch.customerUseCase.GetFiltered(*params.Query)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(presenter.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}
	customerList := []presenter.FilteredCustomer{}
	for _, v := range customers {
		customerList = append(customerList, presenter.FilteredCustomer{
			Id:   v.ID(),
			Name: v.Name(),
		})
	}

	c.JSON(http.StatusOK, presenter.FilteredCustomersListResponse{
		ApiVersion: api.Version,
		Items:      customerList,
	})
}

func (ch *CustomerHandler) GetAllCustomersCount(c *gin.Context) {
	count, err := ch.customerUseCase.GetAllCount()
	if err != nil {
		logger.Error(err.Error())
		c.JSON(presenter.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, presenter.AllCustomersCountResponse{
		ApiVersion: api.Version,
		Data: presenter.AllCustomerCount{
			Total: int(count),
		},
	})
}
