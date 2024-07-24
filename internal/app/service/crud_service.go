package service

import (
	"github.com/balireddypraveen/allen/internal/app/dto"
	"github.com/balireddypraveen/allen/internal/app/models"
	"github.com/balireddypraveen/allen/internal/app/repo"
	"github.com/balireddypraveen/allen/internal/pkg/context"
	"github.com/google/uuid"
)

type ICrudService interface {
	CreateOrder(reqCtx context.ReqCtx, createOrderRequest dto.CreateOrderRequest) (*models.Order, error)
	ReadOrders(reqCtx context.ReqCtx, getOrdersDTO dto.GetOrdersRequest) ([]models.Order, error)
	UpdateOrder(reqCtx context.ReqCtx, createOrderRequest dto.CreateOrderRequest) (*models.Order, error)
	DeleteOrder(reqCtx context.ReqCtx, cancelOrderRequest dto.CancelOrderRequest) error
}

type CrudService struct {
	crudRepo *repo.CrudRepo
}

func NewCrudService(crudRepo *repo.CrudRepo) *CrudService {
	return &CrudService{
		crudRepo: crudRepo,
	}
}

func (c CrudService) CreateOrder(reqCtx context.ReqCtx, createOrderRequest dto.CreateOrderRequest) (*models.Order, error) {
	orderId, _ := uuid.NewUUID()
	order := models.Order{
		OrderID:   orderId,
		Status:    "INIT",
		UserId:    createOrderRequest.UserId,
	}
	order, err := c.crudRepo.CreateOrder(reqCtx, order)
	return &order,err
}

func (c CrudService) ReadOrders(reqCtx context.ReqCtx, getOrdersDTO dto.GetOrdersRequest) ([]models.Order, error) {
	panic("implement me")
}

func (c CrudService) UpdateOrder(reqCtx context.ReqCtx, createOrderRequest dto.CreateOrderRequest) (*models.Order, error) {
	panic("implement me")
}

func (c CrudService) DeleteOrder(reqCtx context.ReqCtx, cancelOrderRequest dto.CancelOrderRequest) error {
	panic("implement me")
}
