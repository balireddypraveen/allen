package rest

import (
	"github.com/balireddypraveen/allen/internal/common/constants"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(routerGroup *gin.RouterGroup) *gin.RouterGroup {
	container := NewContainer()
	v1 := routerGroup.Group(constants.V1)

	externalRoutes(v1, &container)
	apiRoutes(v1, &container)

	return v1
}

func healthCheck(context *gin.Context) {
	resp := map[string]string{
		"deployed_image": os.Getenv("DEPLOYED_IMAGE_TAG"),
	}
	context.JSON(http.StatusOK, resp)
}

func externalRoutes(routerGroup *gin.RouterGroup, container *Container) {
	routerGroup.GET(constants.TestHealth, healthCheck)
}

func apiRoutes(routerGroup *gin.RouterGroup, container *Container) {
	routerGroup.POST(constants.CreateOrder, container.apiController.Create)
}
