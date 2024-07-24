package repo

import (
	"fmt"
	"github.com/balireddypraveen/allen/internal/app/repo/base_repo"
	customContext "github.com/balireddypraveen/allen/internal/pkg/context"

	"github.com/balireddypraveen/allen/internal/app/models"
	"github.com/google/uuid"
)

type IDealRepository interface {
	CreateDeal(reqCtx customContext.ReqCtx, order models.Deal) (models.Deal, error)
	GetDealById(reqCtx customContext.ReqCtx, dealId uuid.UUID) (*models.Deal, error)
	UpdateDealsWhere(ctx customContext.ReqCtx, fields interface{}, conditions interface{}) error
	EndDealById(reqCtx customContext.ReqCtx, dealId uuid.UUID) error
}

type DealRepo struct {
	BaseRepo  base_repo.IBaseRepo
	tableName string
}

func NewDealRepo(baseRepo *base_repo.BaseRepo) *DealRepo {
	return &DealRepo{
		BaseRepo:  baseRepo,
		tableName: "deals",
	}
}

func (c DealRepo) CreateDeal(reqCtx customContext.ReqCtx, deal models.Deal) (models.Deal, error) {
	err := c.BaseRepo.Create(reqCtx, &deal, c.tableName)
	if err != nil {
		reqCtx.Log.Errorf("error creating order in db, order_id:- %v, err is %v", deal.DealId, err.Error())
		return deal, err
	}
	return deal, nil
}

func (c DealRepo) UpdateDealsWhere(ctx customContext.ReqCtx, fields interface{}, conditions interface{}) error {
	panic("implement me")
}

func (c DealRepo) EndDealById(reqCtx customContext.ReqCtx, dealId uuid.UUID) error {
	panic("implement me")
}

func (c DealRepo) GetDealById(reqCtx customContext.ReqCtx, dealId uuid.UUID) (*models.Deal, error) {
	var deal *models.Deal
	condition := fmt.Sprintf("deal_id = '%s'", dealId.String())
	err := c.BaseRepo.GetRecordsByCondition(reqCtx, &deal, condition, c.tableName)
	return deal, err
}
