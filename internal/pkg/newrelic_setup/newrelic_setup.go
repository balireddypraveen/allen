package newrelic_setup

import (
	"context"
	"github.com/balireddypraveen/allen/internal/common/configs"
	"github.com/balireddypraveen/allen/internal/pkg/logger"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/spf13/viper"
	"sync"
)

var nrApp *newrelic.Application
var once sync.Once

func IsNewrelicEnable() bool {
	return viper.GetBool(configs.VKEYS_NEWRELIC_ENABLED)
}

func newrelicInit(applicationName string) (*newrelic.Application, error) {
	var err error
	once.Do(func() {
		appName := applicationName + "-" + viper.GetString(configs.VKEYS_HOST_TYPE)
		license := viper.GetString(configs.VKEYS_NEWRELIC_LICENSE)
		nrApp, err = newrelic.NewApplication(
			newrelic.ConfigAppName(appName),
			newrelic.ConfigLicense(license),
			newrelic.ConfigDistributedTracerEnabled(true),
		)
	})
	return nrApp, err
}

func GetNewRelicApp(applicationName string) *newrelic.Application {
	log := logger.GetLogger()
	isNewRelicEnable := IsNewrelicEnable()
	if isNewRelicEnable {
		//log.Debug("starting new relic!")
		app, err := newrelicInit(applicationName)
		if nil != err {
			log.Info("error creating app (invalid newrelic config):", err)
			return nil
		}
		return app
	}
	return nil
}

// NewRelicForRedisStartSegment setupNewRelicSegment start a datastore segment for the redis call
// it starts the segment and return the End function which can be deferred
func NewRelicForRedisStartSegment(ctx context.Context, operation string, key string) func() {
	txn := newrelic.FromContext(ctx)
	newRelicSegment := newrelic.DatastoreSegment{
		Product:            newrelic.DatastoreRedis,
		Host:               viper.GetString(configs.VKEYS_REDIS_CLUSTERS_HOST_URL),
		Operation:          operation,
		ParameterizedQuery: key,
	}
	newRelicSegment.StartTime = txn.StartSegmentNow()
	return newRelicSegment.End

}
