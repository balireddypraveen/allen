package controller

import (
	"github.com/balireddypraveen/allen/internal/app/dto"
	"github.com/balireddypraveen/allen/internal/app/service"
	customContext "github.com/balireddypraveen/allen/internal/pkg/context"
	"github.com/balireddypraveen/allen/internal/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ICrudController interface {
	Create(ctx *gin.Context)
	Read(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}


type CrudController struct {
	crudService service.ICrudService
}

func NewCrudController(crudService service.ICrudService) *CrudController {
	return &CrudController{
		crudService: crudService,
	}
}

func (c CrudController) Create(ctx *gin.Context) {
	rCtx := customContext.GetRequestContext(ctx)
	log := rCtx.Log

	var createOrderDTO dto.CreateOrderRequest
	if err := ctx.ShouldBindJSON(&createOrderDTO); err != nil {
		log.Errorf("create order request invalid, error is %+v", err.Error())
		response.Error(ctx, rCtx, http.StatusBadRequest, response.InvalidInput, err.Error())
		return
	}

	createOrderResponse, err := c.crudService.CreateOrder(rCtx, createOrderDTO)
	if err != nil {
		log.Errorf("Error in Create Order %+v :err is  %+v", createOrderDTO, err)
		response.Error(ctx, rCtx, http.StatusInternalServerError, response.ErrorWhileProcessing, err.Error())
		return
	}

	response.Success(ctx, rCtx, http.StatusOK, createOrderResponse)
	return
}

func (c CrudController) Read(ctx *gin.Context) {
	panic("implement me")
}

func (c CrudController) Update(ctx *gin.Context) {
	panic("implement me")
}

func (c CrudController) Delete(ctx *gin.Context) {
	panic("implement me")
}