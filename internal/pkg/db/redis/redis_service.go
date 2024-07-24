package redis

import (
	"github.com/balireddypraveen/allen/internal/pkg/context"
	"time"
)

type IRedis interface {
	Get(rCtx context.ReqCtx, key string) (map[string]interface{}, error)
	GetUnmarshalled(rCtx context.ReqCtx, key string, data interface{}) error
	Set(rCtx context.ReqCtx, key string, value interface{}, timeout time.Duration) error
	AcquireLock(rCtx context.ReqCtx, lockingKey string, expiry time.Duration) (mutex interface{}, err error)
	ReleaseLock(rCtx context.ReqCtx, mutex interface{}) (err error)
}
