package rest

import (
	"github.com/balireddypraveen/allen/internal/common/configs"
	"github.com/balireddypraveen/allen/internal/common/constants"
	"github.com/balireddypraveen/allen/internal/pkg/newrelic_setup"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/spf13/viper"
	"math/rand"
	"net/http"
	"time"
)

// BuildServer adds all the configurations and returns the pointer to the server
func BuildServer() *gin.Engine {
	server := gin.New()
	//server.Use(middleware.QueryLogMiddleware())
	//server.Use(CustomRecovery())
	server.Use(cors.New(cors.Config{
		AllowOrigins:     viper.GetStringSlice(configs.VKEYS_CORS_ORIGINS),
		AllowMethods:     viper.GetStringSlice(configs.VKEYS_ALLOW_METHODS),
		AllowHeaders:     viper.GetStringSlice(configs.VKEYS_SERVER_ALLOW_HEADERS),
		ExposeHeaders:    viper.GetStringSlice(configs.VKEYS_EXPOSED_HEADERS),
		AllowCredentials: true,
	}))

	// add new relic to all routes, it attaches the transaction context to gin context and is picked up
	//server.Use(middleware.RequestRequirements())
	server.Use(nrgin.Middleware(newrelic_setup.GetNewRelicApp(viper.GetString(configs.VKEYS_NEWRELIC_APP_NAME))))
	RegisterRoutes(server.Group(constants.ApiPath))

	// initialize global pseudo random generator
	rand.Seed(time.Now().Unix())

	return server
}

// HttpBuildServer for solving timeout issue - http server using gin router
func HttpBuildServer(addr string) *http.Server {
	server := BuildServer()
	s := &http.Server{
		Addr:         addr,
		Handler:      server,
		ReadTimeout:  time.Duration(viper.GetInt(configs.VKEYS_READ_WRITE_TIMEOUT_SERVER)) * time.Second,
		WriteTimeout: time.Duration(viper.GetInt(configs.VKEYS_READ_WRITE_TIMEOUT_SERVER)) * time.Second,
	}
	s.IdleTimeout = time.Duration(viper.GetInt(configs.VKEYS_IDLE_TIMEOUT_SERVER)) * time.Minute

	return s
}
