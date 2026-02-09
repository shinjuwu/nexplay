package redis

import (
	"backend/pkg/config"
	"context"
	"errors"

	"sync"
	"time"

	"github.com/go-redis/redis/v9"
	"go.uber.org/zap"
)

type RedisCliectV9 struct {
	// redis db object map [dbIdx, *redis.Client]
	ctx          context.Context
	redis_cliect *sync.Map
}

func NewRedisCliectV9(multiLogger *zap.Logger, config config.Config) *RedisCliectV9 {

	rc := &RedisCliectV9{
		ctx:          context.Background(),
		redis_cliect: &sync.Map{},
	}

	for i := 0; i < len(config.GetRedis().DbIndex); i++ {
		err := rc.CreateRedisObject(config.GetRedis().DbIndex[i], config.GetRedis().Address, config.GetRedis().Password)
		if err != nil {
			multiLogger.Fatal("Redis database ping error", zap.Error(err))
		}
	}
	multiLogger.Info("Redis open index", zap.Ints("index array", config.GetRedis().DbIndex))

	return rc
}

func NewRedisCliectV9Test() *RedisCliectV9 {
	return &RedisCliectV9{
		ctx:          context.Background(),
		redis_cliect: &sync.Map{},
	}
}

// 創建一個redis 物件
func (r *RedisCliectV9) CreateRedisObject(dbIdx int, ipAddress string, password string) error {
	// r := redisModule.Init("127.0.0.1:6379", "", dbIdx)

	rdb := redis.NewClient(&redis.Options{
		Addr:     ipAddress,
		Password: password, // no password set
		DB:       dbIdx,    // use default DB
	})

	// ping to check create success
	if _, err := rdb.Ping(r.ctx).Result(); err != nil {
		return err
	} else {
		r.redis_cliect.Store(dbIdx, rdb)
		return nil
	}
}

///////////////////////////////////////////////////////////////////////////////
/*
one key count value
*/
///////////////////////////////////////////////////////////////////////////////

// 依照key值,將計數加1
func (r *RedisCliectV9) Incr(dbIdx int, key string) error {
	if v, ok := r.redis_cliect.Load(dbIdx); ok {
		client := v.(*redis.Client)
		return client.Incr(r.ctx, key).Err()
	} else {
		return errors.New("no register redis object")
	}
}

// 依照key值,將計數增加指定數值
func (r *RedisCliectV9) IncrBy(dbIdx int, key string, value int64) error {
	if v, ok := r.redis_cliect.Load(dbIdx); ok {
		client := v.(*redis.Client)
		return client.IncrBy(r.ctx, key, value).Err()
	} else {
		return errors.New("no register redis object")
	}
}

// 依照key值,將計數減1
func (r *RedisCliectV9) Decr(dbIdx int, key string) error {
	if v, ok := r.redis_cliect.Load(dbIdx); ok {
		client := v.(*redis.Client)
		return client.Decr(r.ctx, key).Err()
	} else {
		return errors.New("no register redis object")
	}
}

// 依照key值,將計數減少指定數值
func (r *RedisCliectV9) DecrBy(dbIdx int, key string, value int64) error {
	if v, ok := r.redis_cliect.Load(dbIdx); ok {
		client := v.(*redis.Client)
		return client.DecrBy(r.ctx, key, value).Err()
	} else {
		return errors.New("no register redis object")
	}
}

// 依照key值,將對應的值直接設定為指定的值
func (r *RedisCliectV9) GetSet(dbIdx int, key string, value interface{}) error {
	if v, ok := r.redis_cliect.Load(dbIdx); ok {
		client := v.(*redis.Client)
		return client.GetSet(r.ctx, key, value).Err()
	} else {
		return errors.New("no register redis object")
	}
}

///////////////////////////////////////////////////////////////////////////////
/*
one key one value
*/
///////////////////////////////////////////////////////////////////////////////

/*
Redis `SET key value [expiration]` command.

Use expiration for `SETEX`-like behavior.
Zero expiration means the key has no expiration time.
*/
func (r *RedisCliectV9) StoreValue(dbIdx int, key string, value interface{}, expiration time.Duration) error {
	if v, ok := r.redis_cliect.Load(dbIdx); ok {
		client := v.(*redis.Client)
		return client.Set(r.ctx, key, value, expiration).Err()
	} else {
		return errors.New("no register redis object")
	}
}

// Redis `GET key` command. It returns redis.Nil error when key does not exist..
func (r *RedisCliectV9) LoadValue(dbIdx int, key string) (string, error) {
	if v, ok := r.redis_cliect.Load(dbIdx); ok {
		client := v.(*redis.Client)
		return client.Get(r.ctx, key).Result()
	} else {
		return "", errors.New("no register redis object")
	}
}

/*
確認某個key值下是否有資料存在
*/
func (r *RedisCliectV9) CheckValue(dbIdx int, key string) (int64, error) {
	if v, ok := r.redis_cliect.Load(dbIdx); ok {
		client := v.(*redis.Client)
		return client.Exists(r.ctx, key).Result()
	} else {
		return 0, errors.New("no register redis object")
	}
}

/*
刪除某一個key值
*/
func (r *RedisCliectV9) DeleteValue(dbIdx int, key string) error {
	if v, ok := r.redis_cliect.Load(dbIdx); ok {
		client := v.(*redis.Client)
		return client.Del(r.ctx, key).Err()
	} else {
		return errors.New("no register redis object")
	}
}

///////////////////////////////////////////////////////////////////////////////
/*
one hash have more [key ,value]

Sets field in the hash stored at key to value. If key does not exist, a new key holding a hash is created. If field already exists in the hash, it is overwritten.

sample format:
hash: "XXX"
key: "test"
value: "123456"
*/
///////////////////////////////////////////////////////////////////////////////

/*
創建/覆蓋某個hash下的某一個或多個field
*/
func (r *RedisCliectV9) StoreHValue(dbIdx int, hash string, args ...string) error {
	if v, ok := r.redis_cliect.Load(dbIdx); ok {
		client := v.(*redis.Client)
		return client.HSet(r.ctx, hash, args).Err()
	} else {
		return errors.New("no register redis object")
	}
}

/*
取得某個hash下的某一個field
Returns the value associated with field in the hash stored at key.
*/
func (r *RedisCliectV9) LoadHValue(dbIdx int, hash string, field string) (string, error) {
	if v, ok := r.redis_cliect.Load(dbIdx); ok {
		client := v.(*redis.Client)
		return client.HGet(r.ctx, hash, field).Result()
	} else {
		return "", errors.New("no register redis object")
	}
}

/*
取得某個hash下的全部field
Returns the value associated with field in the hash stored at key.
*/
func (r *RedisCliectV9) LoadHAllValue(dbIdx int, hash string) (map[string]string, error) {
	if v, ok := r.redis_cliect.Load(dbIdx); ok {
		client := v.(*redis.Client)
		return client.HGetAll(r.ctx, hash).Result()
	} else {
		return nil, errors.New("no register redis object")
	}
}

// func (r *RedisCliectV9) LoadAllHValue(dbIdx int, keys []string, field string) ([]redis.Cmder, error) {
// 	if v, ok := r.redis_cliect.Load(dbIdx); ok {
// 		client := v.(*redis.Client)
// 		pipe := client.TxPipeline()
// 		for _, key := range keys {
// 			pipe.HGet(r.ctx, key, field)
// 		}
// 		return pipe.Exec(r.ctx)
// 	} else {
// 		return nil, errors.New("no register redis object")
// 	}
// }

/*
確認某個hash下的某一個field是否存在
Returns if field is an existing field in the hash stored at key.
*/
func (r *RedisCliectV9) CheckHValue(dbIdx int, hash string, field string) (bool, error) {
	if v, ok := r.redis_cliect.Load(dbIdx); ok {
		client := v.(*redis.Client)
		return client.HExists(r.ctx, hash, field).Result()
	} else {
		return false, errors.New("no register redis object")
	}
}

/*
刪除某個hash下的某一個field
*/
func (r *RedisCliectV9) DeleteHValue(dbIdx int, hash string, field string) error {
	if v, ok := r.redis_cliect.Load(dbIdx); ok {
		client := v.(*redis.Client)
		return client.HDel(r.ctx, hash, field).Err()
	} else {
		return errors.New("no register redis object")
	}
}

func (r *RedisCliectV9) StoreLValue(dbIdx int, key string, values []interface{}) (int64, error) {
	if v, ok := r.redis_cliect.Load(dbIdx); ok {
		client := v.(*redis.Client)
		return client.RPush(r.ctx, key, values...).Result()
	} else {
		return -1, errors.New("no register redis object")
	}
}

// func (r *RedisCliectV9) LoadAllLALLValue(dbIdx int, keys []string) ([]redis.Cmder, error) {
// 	if v, ok := r.redis_cliect.Load(dbIdx); ok {
// 		client := v.(*redis.Client)
// 		pipe := client.TxPipeline()
// 		for _, key := range keys {
// 			pipe.LRange(r.ctx, key, 0, -1)
// 		}
// 		return pipe.Exec(r.ctx)
// 	} else {
// 		return nil, errors.New("no register redis object")
// 	}
// }

/* Redis `flushdb` command.

刪除指定redis DB內的儲存值
*/
func (r *RedisCliectV9) FlushDB(dbIdx int) error {
	if v, ok := r.redis_cliect.Load(dbIdx); ok {
		client := v.(*redis.Client)
		return client.FlushDB(r.ctx).Err()
	} else {
		return errors.New("no register redis object")
	}
}

/* Redis `flushdb` command.

刪除指定redis DB內的儲存值
*/
func (r *RedisCliectV9) Del(dbIdx int, key string) error {
	if v, ok := r.redis_cliect.Load(dbIdx); ok {
		client := v.(*redis.Client)
		newkey := key + "_del"
		client.Rename(r.ctx, key, newkey)
		return client.Del(r.ctx, newkey).Err()
	} else {
		return errors.New("no register redis object")
	}
}

/* Redis `Rename` command.

更改指定redis DB內的 key 值
*/
func (r *RedisCliectV9) Rename(dbIdx int, key, newkey string) error {
	if v, ok := r.redis_cliect.Load(dbIdx); ok {
		client := v.(*redis.Client)
		return client.Rename(r.ctx, key, newkey).Err()
	} else {
		return errors.New("no register redis object")
	}
}
