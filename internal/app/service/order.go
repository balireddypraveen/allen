package service

import (
	"fmt"
	"github.com/balireddypraveen/allen/internal/app/dto"
	"github.com/balireddypraveen/allen/internal/app/models"
	"github.com/balireddypraveen/allen/internal/app/repo"
	"github.com/balireddypraveen/allen/internal/pkg/context"
	"github.com/google/uuid"
	"time"
)

type IOrderService interface {
	CreateOrder(reqCtx context.ReqCtx, req dto.CreateOrderRequest) (*models.Order, error)
}

type OrderService struct {
	orderRepo *repo.OrderRepo
	dealRepo  *repo.DealRepo
}

func NewOrderService(crudRepo *repo.OrderRepo, dealRepo *repo.DealRepo) *OrderService {
	return &OrderService{
		orderRepo: crudRepo,
		dealRepo:  dealRepo,
	}
}

func (c OrderService) CreateOrder(reqCtx context.ReqCtx, req dto.CreateOrderRequest) (*models.Order, error) {
	// Get deal
	deal, err := c.dealRepo.GetDealById(reqCtx, req.DealId)
	if deal == nil {
		return nil, fmt.Errorf("oops! This deal doesn't exist")
	}
	if err != nil {
		return nil, err
	}

	// check if deal is over or it is claimed max times already
	if deal.EndTime.Before(time.Now()) {
		return nil, fmt.Errorf("oops! Deal time over")
	}

	if deal.MaxQuantity <= 0 {
		return nil, fmt.Errorf("oops! Deal claimed already for the maximum times, better luck next time")
	}

	// check if the user has already claimed the deal
	orders, err := c.orderRepo.GetOrdersByUserId(reqCtx, req.UserId)
	if err != nil {
		return nil, err
	}

	if len(orders) > 0 {
		return nil, fmt.Errorf("you have already claimed the deal, check for other deals")
	}

	// All validations passed! Create order now
	orderId, _ := uuid.NewUUID()
	order := models.Order{
		OrderID: orderId,
		UserId:  req.UserId,
		DealId:  req.DealId,
	}
	orderResp, err := c.orderRepo.CreateOrder(reqCtx, order)
	if err != nil {
		return nil, err
	}

	// since the deal is claimed by this user, update the max quantity remaining
	deal.MaxQuantity -= 1
	condition := fmt.Sprintf("deal_id = '%s'", deal.DealId)

	err = c.dealRepo.BaseRepo.UpdateWhere(reqCtx, &deal, "deals", condition)
	if err != nil {
		return nil, err
	}

	return orderResp, err
}
