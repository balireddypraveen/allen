package redis

import (
	"encoding/json"
	"fmt"
	"github.com/balireddypraveen/allen/internal/pkg/context"
	"time"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
)

const RedisPrefix = "{allen}:"

type RedisService struct {
	Redis Redis
}

func NewRedisService() *RedisService {
	return &RedisService{Redis: *Client}
}

func (redisService *RedisService) Get(rCtx context.ReqCtx, key string) (map[string]interface{}, error) {
	if rCtx.NewRelicTxn != nil {
		seg := rCtx.NewRelicTxn.StartSegment("Redis get")
		defer seg.End()
	}

	redisKey := RedisPrefix + key
	val, err := redisService.Redis.Get(rCtx, redisKey)
	if err != nil {
		//rCtx.Log.Errorf("Failed to get element from redis key:- %s. Error :- %s", key, err.Error())
		return nil, err
	}

	var data map[string]interface{}
	err = json.Unmarshal([]byte(val), &data)
	if err != nil {
		rCtx.Log.Errorf("Failed to get unmarshal redis element. Redis key:- %s. Error :- %s", key, err.Error())
		return nil, err
	}

	return data, nil
}

func (redisService *RedisService) GetUnmarshalled(rCtx context.ReqCtx, key string, data interface{}) error {
	if rCtx.NewRelicTxn != nil {
		seg := rCtx.NewRelicTxn.StartSegment("Redis get unmarshalled")
		defer seg.End()
	}

	redisKey := RedisPrefix + key
	val, err := redisService.Redis.GetNew(rCtx, redisKey)
	if err != nil {
		// rCtx.Log.InfofCf("Failed to get element from redis key:- %s. Error :- %s", key, err.Error())
		return err
	}

	err = json.Unmarshal([]byte(val), &data)
	if err != nil {
		rCtx.Log.Errorf("Failed to get unmarshal redis element. Redis key:- %s. Error :- %s", key, err.Error())
		return err
	}

	return nil
}

func (redisService *RedisService) Set(rCtx context.ReqCtx, key string, value interface{}, timeout time.Duration) error {
	if rCtx.NewRelicTxn != nil {
		seg := rCtx.NewRelicTxn.StartSegment("Redis set")
		defer seg.End()
	}

	redisKey := RedisPrefix + key
	valueStr, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return redisService.Redis.SetNew(rCtx, redisKey, string(valueStr), timeout)
}

func (redisService *RedisService) AcquireLock(rCtx context.ReqCtx, lockingKey string, expiry time.Duration) (mutex interface{}, err error) {
	if rCtx.NewRelicTxn != nil {
		seg := rCtx.NewRelicTxn.StartSegment("Redis acquire lock")
		defer seg.End()
	}

	log := rCtx.Log
	pool := goredis.NewPool(redisService.Redis.redis)
	rs := redsync.New(pool)
	mutex = rs.NewMutex(lockingKey, redsync.WithExpiry(expiry))
	log.Debugf("AcquireLock mutex start: (%v)", mutex)
	// Obtain a lock for our given mutex. After this is successful, no one else
	// can obtain the same lock (the same mutex name) until we unlock it.
	err = mutex.(*redsync.Mutex).Lock()
	if err != nil {
		log.Errorf("AcquireLock mutex failed: (%v)", mutex)
		return nil, err
	}
	log.Debugf("AcquireLock mutex end: (%v)", mutex)
	return
}

func (redisService *RedisService) ReleaseLock(rCtx context.ReqCtx, mutex interface{}) (err error) {
	if rCtx.NewRelicTxn != nil {
		seg := rCtx.NewRelicTxn.StartSegment("Redis release lock")
		defer seg.End()
	}

	log := rCtx.Log

	log.Debugf("ReleaseLock mutex start: (%v)", mutex)
	ok, err := mutex.(*redsync.Mutex).Unlock()
	if !ok {
		return fmt.Errorf("mutex unlocking is not ok: %v", err)
	}
	log.Debugf("ReleaseLock mutex end: (%v)", mutex)
	return
}
