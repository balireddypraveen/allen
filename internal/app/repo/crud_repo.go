package repo

import (
	"github.com/balireddypraveen/allen/internal/app/dto"
	"github.com/balireddypraveen/allen/internal/app/repo/base_repo"
	customContext "github.com/balireddypraveen/allen/internal/pkg/context"

	"github.com/balireddypraveen/allen/internal/app/models"
	"github.com/google/uuid"
)

type ICrudRepository interface {
	CreateOrder(reqCtx customContext.ReqCtx, order models.Order) (models.Order, error)
	GetOrders(reqCtx customContext.ReqCtx, getOrdersDTO dto.GetOrdersRequest) ([]models.Order, error)
	UpdateOrdersWhere(ctx customContext.ReqCtx, fields interface{}, conditions interface{}) error
	CancelOrderById(reqCtx customContext.ReqCtx, orderId uuid.UUID) error

	// CreateOrderWithTxn helpers methods for concurrency and locking
	//CreateOrderWithTxn(reqCtx customContext.ReqCtx, txn *gorm.DB, order models.Order) (models.Order, error)
	//GetOrderWithTxnSkipLock(reqCtx customContext.ReqCtx, txn *gorm.DB, orderId uuid.UUID) (models.Order, error)
	//
	//AcquireLock(reqCtx customContext.ReqCtx, id uuid.UUID) error
	//ReleaseLock(reqCtx customContext.ReqCtx, id uuid.UUID) error
	//
	//GetTransaction() *gorm.DB
	//Rollback(rCtx customContext.ReqCtx, tx *gorm.DB)
}

type CrudRepo struct {
	baseRepo base_repo.IBaseRepo
	tableName string
}

func NewCrudRepo(baseRepo *base_repo.BaseRepo) *CrudRepo {
	return &CrudRepo{
		baseRepo: baseRepo,
		tableName: "orders",
	}
}

func (c CrudRepo) CreateOrder(reqCtx customContext.ReqCtx, order models.Order) (models.Order, error) {
	err := c.baseRepo.Create(reqCtx, &order, c.tableName)
	if err != nil {
		reqCtx.Log.Errorf("error creating order in db, order_id:- %v, err is %v", order.OrderID, err.Error())
		return order, err
	}
	return order, nil
}

func (c CrudRepo) GetOrders(reqCtx customContext.ReqCtx, getOrdersDTO dto.GetOrdersRequest) ([]models.Order, error) {
	panic("implement me")
}

func (c CrudRepo) UpdateOrdersWhere(ctx customContext.ReqCtx, fields interface{}, conditions interface{}) error {
	panic("implement me")
}

func (c CrudRepo) CancelOrderById(reqCtx customContext.ReqCtx, orderId uuid.UUID) error {
	panic("implement me")
}


