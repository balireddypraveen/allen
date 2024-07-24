package controller

import (
	"github.com/balireddypraveen/allen/internal/app/dto"
	"github.com/balireddypraveen/allen/internal/app/service"
	customContext "github.com/balireddypraveen/allen/internal/pkg/context"
	"github.com/balireddypraveen/allen/internal/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IOrderController interface {
	CreateOrder(ctx *gin.Context)
}

type OrderController struct {
	orderService service.IOrderService
}

func NewOrderController(orderService service.IOrderService) *OrderController {
	return &OrderController{
		orderService: orderService,
	}
}

func (c OrderController) CreateOrder(ctx *gin.Context) {
	rCtx := customContext.GetRequestContext(ctx)
	log := rCtx.Log

	var req dto.CreateOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Errorf("create deal request invalid, error is %+v", err.Error())
		response.Error(ctx, rCtx, http.StatusBadRequest, response.InvalidInput, err.Error())
		return
	}

	createOrderResponse, err := c.orderService.CreateOrder(rCtx, req)
	if err != nil {
		log.Errorf("Error in CreateOrder Order %+v :err is  %+v", req, err)
		response.Error(ctx, rCtx, http.StatusInternalServerError, response.ErrorWhileProcessing, err.Error())
		return
	}

	response.Success(ctx, rCtx, http.StatusOK, createOrderResponse)
	return
}
