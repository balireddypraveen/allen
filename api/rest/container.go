package rest

import (
	"github.com/balireddypraveen/allen/internal/app/controller"
	"github.com/balireddypraveen/allen/internal/app/repo"
	"github.com/balireddypraveen/allen/internal/app/repo/base_repo"
	"github.com/balireddypraveen/allen/internal/app/service"
	"github.com/balireddypraveen/allen/internal/pkg/db/postgres"
)

type Container struct {
	apiController controller.ICrudController
}

func NewContainer() Container {

	//Base Repo
	baseRepo := base_repo.NewBaseRepo()
	baseRepo.SetDb(postgres.GetDBWithoutContext())

	//Repository
	crudRepo := repo.NewCrudRepo(baseRepo)

	//Service
	crudService := service.NewCrudService(crudRepo)

	//Controller
	apiController := controller.NewCrudController(crudService)

	return Container{
		apiController: apiController,
	}
}
