package redis

import (
	"backend/pkg/config"
	"time"

	"go.uber.org/zap"
)

/*
If you are using Redis 6, install go-redis/v8:

If you are using Redis 7, install go-redis/v9:
*/
type IRedisCliect interface {
	CreateRedisObject(dbIdx int, ipAddress string, password string) error
	Incr(dbIdx int, key string) error
	IncrBy(dbIdx int, key string, value int64) error
	Decr(dbIdx int, key string) error
	DecrBy(dbIdx int, key string, value int64) error
	GetSet(dbIdx int, key string, value interface{}) error
	StoreValue(dbIdx int, key string, value interface{}, expiration time.Duration) error
	LoadValue(dbIdx int, key string) (string, error)
	CheckValue(dbIdx int, key string) (int64, error)
	DeleteValue(dbIdx int, key string) error
	StoreHValue(dbIdx int, hash string, args ...string) error
	LoadHValue(dbIdx int, hash string, field string) (string, error)
	LoadHAllValue(dbIdx int, hash string) (map[string]string, error)
	// LoadAllHValue(dbIdx int, keys []string, field string) ([]redis.Cmder, error)
	CheckHValue(dbIdx int, hash string, field string) (bool, error)
	DeleteHValue(dbIdx int, hash string, field string) error
	StoreLValue(dbIdx int, key string, values []interface{}) (int64, error)
	// LoadAllLALLValue(dbIdx int, keys []string) ([]redis.Cmder, error)
	FlushDB(dbIdx int) error
	Del(dbIdx int, key string) error
	Rename(dbIdx int, key, newkey string) error
}

func NewRedisCliect(multiLogger *zap.Logger, config config.Config) IRedisCliect {

	rdb := NewRedisCliectV9(multiLogger, config)

	return rdb
}
