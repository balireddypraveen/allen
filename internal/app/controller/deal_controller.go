package controller

import (
	"github.com/balireddypraveen/allen/internal/app/dto"
	"github.com/balireddypraveen/allen/internal/app/service"
	customContext "github.com/balireddypraveen/allen/internal/pkg/context"
	"github.com/balireddypraveen/allen/internal/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IDealController interface {
	CreateDeal(ctx *gin.Context)
	UpdateDeal(ctx *gin.Context)
	EndDeal(ctx *gin.Context)
}

type DealController struct {
	dealService service.IDealService
}

func NewDealController(dealService service.IDealService) *DealController {
	return &DealController{
		dealService: dealService,
	}
}

func (c DealController) CreateDeal(ctx *gin.Context) {
	rCtx := customContext.GetRequestContext(ctx)
	log := rCtx.Log

	var createDealRequest dto.CreateDealRequest
	if err := ctx.ShouldBindJSON(&createDealRequest); err != nil {
		log.Errorf("create deal request invalid, error is %+v", err.Error())
		response.Error(ctx, rCtx, http.StatusBadRequest, response.InvalidInput, err.Error())
		return
	}

	createOrderResponse, err := c.dealService.CreateDeal(rCtx, createDealRequest)
	if err != nil {
		log.Errorf("Error in CreateOrder Order %+v :err is  %+v", createDealRequest, err)
		response.Error(ctx, rCtx, http.StatusInternalServerError, response.ErrorWhileProcessing, err.Error())
		return
	}

	response.Success(ctx, rCtx, http.StatusOK, createOrderResponse)
	return
}

func (c DealController) UpdateDeal(ctx *gin.Context) {
	rCtx := customContext.GetRequestContext(ctx)
	log := rCtx.Log

	var req dto.UpdateDealRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Errorf("update deal request invalid, error is %+v", err.Error())
		response.Error(ctx, rCtx, http.StatusBadRequest, response.InvalidInput, err.Error())
		return
	}

	createOrderResponse, err := c.dealService.UpdateDeal(rCtx, req)
	if err != nil {
		log.Errorf("Error in UpdateDeal  %+v :err is  %+v", req, err)
		response.Error(ctx, rCtx, http.StatusInternalServerError, response.ErrorWhileProcessing, err.Error())
		return
	}

	response.Success(ctx, rCtx, http.StatusOK, createOrderResponse)
	return
}

func (c DealController) EndDeal(ctx *gin.Context) {
	panic("implement me")
}
