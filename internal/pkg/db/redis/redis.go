package redis

import (
	"context"
	"github.com/balireddypraveen/allen/internal/common/configs"
	customContext "github.com/balireddypraveen/allen/internal/pkg/context"
	log "github.com/sirupsen/logrus"
	"time"

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var Client = &Redis{}

type Redis struct {
	redis *redis.Client
	ctx   customContext.ReqCtx
}

func InitRedis() {
	ctx := context.TODO()

	host := viper.GetString(configs.VKEYS_REDIS_CLUSTERS_HOST_URL)

	GlobalRedisClient := redis.NewClient(&redis.Options{
		Addr: host,
	})

	_, err := GlobalRedisClient.Ping(ctx).Result()
	if err != nil {
		log.Error("[CRITICAL] failed to initialize Redis error:-" + err.Error())
		return
	}
	Client.redis = GlobalRedisClient

	log.Info("redis initialized successfully")
}

// GetNew CreateOrder new function
func (r *Redis) GetNew(rCtx customContext.ReqCtx, key string) (string, error) {
	// defer setupNewRelicSegment(&ctx, "get", key)()
	log := rCtx.Log
	// log.Infof("In redis.go: ", key, rCtx)
	result, err := Client.redis.Get(rCtx, key).Result()
	if err != nil {
		if err != redis.Nil {
			log.Errorf("failed while getting redis client error - %v", err.Error())
		}
		return "", err
	}
	// log.Infof("In redis.go again: ", key, rCtx, result)
	return result, nil
}

func (r *Redis) SetNew(rCtx customContext.ReqCtx, key, value string, timeout time.Duration) error {
	// defer setupNewRelicSegment(&ctx, "set", key)()
	err := Client.redis.Set(rCtx, key, value, timeout).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *Redis) Get(ctx customContext.ReqCtx, key string) (string, error) {
	defer setupNewRelicSegment(&ctx, "get", key)()
	result, err := r.redis.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}

func (r *Redis) Set(ctx customContext.ReqCtx, key, value string, timeout time.Duration) error {
	defer setupNewRelicSegment(&ctx, "set", key)()
	err := r.redis.Set(r.ctx, key, value, timeout).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *Redis) MGet(ctx customContext.ReqCtx, keys ...string) ([]string, error) {
	defer setupNewRelicSegment(&ctx, "mget", "")()
	result, err := r.redis.MGet(r.ctx, keys...).Result()
	if err != nil {
		return nil, err
	}
	values := make([]string, len(result))
	for i, val := range result {
		if val == nil {
			values[i] = ""
		} else {
			values[i] = val.(string)
		}
	}
	return values, nil
}

func (r *Redis) HGet(ctx customContext.ReqCtx, key, field string) (string, error) {
	defer setupNewRelicSegment(&ctx, "hget", key)()
	result, err := r.redis.HGet(r.ctx, key, field).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}

func (r *Redis) HSet(ctx customContext.ReqCtx, key string, field string, value interface{}) error {
	defer setupNewRelicSegment(&ctx, "hset", key)()
	err := r.redis.HSet(r.ctx, key, field, value).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *Redis) HMGet(ctx customContext.ReqCtx, key string, fields ...string) (map[string]string, error) {
	defer setupNewRelicSegment(&ctx, "hmget", key)()
	result, err := r.redis.HMGet(r.ctx, key, fields...).Result()
	if err != nil {
		return nil, err
	}

	values := make(map[string]string)
	for i, val := range result {
		if val != nil {
			values[fields[i]] = val.(string)
		}
	}

	return values, nil
}

func (r *Redis) HGetAll(ctx customContext.ReqCtx, key string) (map[string]string, error) {
	defer setupNewRelicSegment(&ctx, "hgetall", key)()
	result, err := r.redis.HGetAll(r.ctx, key).Result()
	if err != nil {
		return nil, err
	}
	values := make(map[string]string)
	for field, val := range result {
		values[field] = val
	}
	return values, nil
}

// setupNewRelicSegment start a datastore segment for the redis call
// it starts the segment and return the End function which can be deferred
func setupNewRelicSegment(rCtx *customContext.ReqCtx, operation string, key string) func() {
	newRelicSegment := newrelic.DatastoreSegment{
		Product:            newrelic.DatastoreRedis,
		Host:               viper.GetString(configs.VKEYS_REDIS_CLUSTERS_HOST_URL),
		Operation:          operation,
		ParameterizedQuery: key,
		RawQuery:           key,
	}
	newRelicSegment.StartTime = rCtx.NewRelicTxn.StartSegmentNow()
	return newRelicSegment.End

}
