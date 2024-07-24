package repo

import (
	"fmt"
	"github.com/balireddypraveen/allen/internal/app/repo/base_repo"
	customContext "github.com/balireddypraveen/allen/internal/pkg/context"

	"github.com/balireddypraveen/allen/internal/app/models"
)

type IOrderRepository interface {
	CreateOrder(reqCtx customContext.ReqCtx, order models.Order) (*models.Order, error)
	GetOrdersByUserId(reqCtx customContext.ReqCtx, userId int) ([]models.Order, error)
}

type OrderRepo struct {
	BaseRepo  base_repo.IBaseRepo
	tableName string
}

func NewOrderRepo(baseRepo *base_repo.BaseRepo) *OrderRepo {
	return &OrderRepo{
		BaseRepo:  baseRepo,
		tableName: "orders",
	}
}

func (c OrderRepo) CreateOrder(reqCtx customContext.ReqCtx, order models.Order) (*models.Order, error) {
	err := c.BaseRepo.Create(reqCtx, &order, c.tableName)
	if err != nil {
		reqCtx.Log.Errorf("error creating order in db, order_id:- %v, err is %v", order.OrderID, err.Error())
		return nil, err
	}
	return &order, nil
}

func (c OrderRepo) GetOrdersByUserId(reqCtx customContext.ReqCtx, userId int) ([]models.Order, error) {
	var order []models.Order
	condition := fmt.Sprintf("user_id = %d", userId)
	err := c.BaseRepo.GetRecordsByCondition(reqCtx, &order, condition, c.tableName)
	return order, err
}
