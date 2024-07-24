package service

import (
	"fmt"
	"github.com/balireddypraveen/allen/internal/app/dto"
	"github.com/balireddypraveen/allen/internal/app/models"
	"github.com/balireddypraveen/allen/internal/app/repo"
	"github.com/balireddypraveen/allen/internal/pkg/common_utils"
	"github.com/balireddypraveen/allen/internal/pkg/context"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"

	//log "github.com/sirupsen/logrus"
	"time"
)

type IDealService interface {
	CreateDeal(reqCtx context.ReqCtx, createDealRequest dto.CreateDealRequest) (*models.Deal, error)
	//GetDeal(reqCtx context.ReqCtx, dealId uuid.UUID) (*models.Deal, error)
	UpdateDeal(reqCtx context.ReqCtx, createOrderRequest dto.UpdateDealRequest) (*models.Deal, error)
	EndDeal(reqCtx context.ReqCtx, cancelOrderRequest dto.EndDealRequest) error
}

type DealService struct {
	dealRepo *repo.DealRepo
}

func NewDealService(dealRepo *repo.DealRepo) *DealService {
	return &DealService{
		dealRepo: dealRepo,
	}
}

func (c DealService) CreateDeal(reqCtx context.ReqCtx, createDealRequest dto.CreateDealRequest) (*models.Deal, error) {
	dealId, _ := uuid.NewUUID()
	dealStartTime, dealEndTime, err := common_utils.GetStartAndEndTimeForDeal()

	deal := models.Deal{
		DealId:      dealId,
		StartTime:   dealStartTime,
		DealName:    createDealRequest.DealName,
		EndTime:     dealEndTime,
		MaxQuantity: createDealRequest.MaxQuantity,
		Enabled:     true,
	}
	deal, err = c.dealRepo.CreateDeal(reqCtx, deal)
	return &deal, err
}

func (c DealService) UpdateDeal(reqCtx context.ReqCtx, req dto.UpdateDealRequest) (*models.Deal, error) {
	deal, err := c.dealRepo.GetDealById(reqCtx, req.DealId)
	log.Info("DEAL is ", deal)
	if deal == nil {
		return nil, fmt.Errorf("sorry, the deal not found")
	}
	if err != nil {
		return nil, err
	}

	// validations
	if deal.EndTime.Before(time.Now()) {
		return nil, fmt.Errorf("sorry, the deal is already closed at %+v", deal.EndTime)
	}
	if deal.MaxQuantity >= req.MaxQuantity {
		return nil, fmt.Errorf("you can't decrease the max quantity. currently it is %v", deal.MaxQuantity)
	} else {
		deal.MaxQuantity = req.MaxQuantity
	}

	if req.ExtendEndTimeByMinutes > 0 {
		*deal.EndTime = deal.EndTime.Add(time.Minute * time.Duration(req.ExtendEndTimeByMinutes))
	}
	c.dealRepo.BaseRepo.Save(reqCtx, deal)
	return deal, nil
}

func (c DealService) EndDeal(reqCtx context.ReqCtx, cancelOrderRequest dto.EndDealRequest) error {
	panic("implement me")
}

//func (c DealService) GetDeal(reqCtx context.ReqCtx, dealId uuid.UUID) (*models.Deal, error)  {
//	dea
//}
