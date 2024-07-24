package main

import (
	"github.com/balireddypraveen/allen/api/rest"
	configs2 "github.com/balireddypraveen/allen/configs"
	"github.com/balireddypraveen/allen/internal/common/configs"
	"github.com/balireddypraveen/allen/internal/common/constants"
	"github.com/balireddypraveen/allen/internal/pkg/db/postgres"
	"github.com/balireddypraveen/allen/internal/pkg/db/redis"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	log.Info("loading config...") // log is initialized after config setup
	configs2.SetupConfig()

	log.Info("initializing postgres db connection...")
	postgres.InitDB()

	log.Info("initializing redis connection...")
	redis.InitRedis()

	log.Info("Setup complete")

}

func getAddr() string {
	addr := constants.ConnectionAddress
	ip := viper.GetString(configs.VKEYS_HOST_IP)
	port := viper.GetString(configs.VKEYS_HOST_PORT)
	if ip != constants.EmptyString || port != constants.EmptyString {
		addr = ip + constants.COLLON + port
	}
	return addr
}

func main() {
	addr := getAddr()
	server := rest.HttpBuildServer(addr)
	log.Info(addr)

	log.Infof("http server is at %v", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Errorf("error while starting server %v", err.Error())
		return
	}
	log.Infof("http server is started")

}
