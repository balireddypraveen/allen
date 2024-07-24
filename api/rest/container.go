package rest

import (
	"github.com/balireddypraveen/allen/internal/app/controller"
	"github.com/balireddypraveen/allen/internal/app/repo"
	"github.com/balireddypraveen/allen/internal/app/repo/base_repo"
	"github.com/balireddypraveen/allen/internal/app/service"
	"github.com/balireddypraveen/allen/internal/pkg/db/postgres"
)

type Container struct {
	apiController   controller.IDealController
	orderController controller.IOrderController
}

func NewContainer() Container {

	//Base Repo
	baseRepo := base_repo.NewBaseRepo()
	baseRepo.SetDb(postgres.GetDBWithoutContext())

	//Repository
	dealRepo := repo.NewDealRepo(baseRepo)
	orderRepo := repo.NewOrderRepo(baseRepo)

	//Service
	dealService := service.NewDealService(dealRepo)
	orderService := service.NewOrderService(orderRepo, dealRepo)

	//Controller
	apiController := controller.NewDealController(dealService)
	orderController := controller.NewOrderController(orderService)

	return Container{
		apiController:   apiController,
		orderController: orderController,
	}
}
