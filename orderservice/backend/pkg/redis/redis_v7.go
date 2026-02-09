package redis

import (
	"backend/pkg/config"
	"errors"

	"sync"
	"time"

	"github.com/go-redis/redis/v7"
	"go.uber.org/zap"
)

type RedisCliect struct {
	// redis db object map [dbIdx, *redis.Client]
	redis_cliect *sync.Map
}

func NewRedisCliectV7(multiLogger *zap.Logger, config config.Config) *RedisCliect {

	rc := &RedisCliect{
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

func NewRedisCliectTest() *RedisCliect {
	return &RedisCliect{
		redis_cliect: &sync.Map{},
	}
}

// 創建一個redis 物件
func (r *RedisCliect) CreateRedisObject(dbIdx int, ipAddress string, password string) error {
	// r := redisModule.Init("127.0.0.1:6379", "", dbIdx)

	rdb := redis.NewClient(&redis.Options{
		Addr:     ipAddress,
		Password: password, // no password set
		DB:       dbIdx,    // use default DB
	})

	// ping to check create success
	if _, err := rdb.Ping().Result(); err != nil {
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
func (r *RedisCliect) Incr(dbIdx int, key string) error {
	if v, ok := r.redis_cliect.Load(dbIdx); ok {
		client := v.(*redis.Client)
		return client.Incr(key).Err()
	} else {
		return errors.New("no register redis object")
	}
}

// 依照key值,將計數增加指定數值
func (r *RedisCliect) IncrBy(dbIdx int, key string, value int64) error {
	if v, ok := r.redis_cliect.Load(dbIdx); ok {
		client := v.(*redis.Client)
		return client.IncrBy(key, value).Err()
	} else {
		return errors.New("no register redis object")
	}
}

// 依照key值,將計數減1
func (r *RedisCliect) Decr(dbIdx int, key string) error {
	if v, ok := r.redis_cliect.Load(dbIdx); ok {
		client := v.(*redis.Client)
		return client.Decr(key).Err()
	} else {
		return errors.New("no register redis object")
	}
}

// 依照key值,將計數減少指定數值
func (r *RedisCliect) DecrBy(dbIdx int, key string, value int64) error {
	if v, ok := r.redis_cliect.Load(dbIdx); ok {
		client := v.(*redis.Client)
		return client.DecrBy(key, value).Err()
	} else {
		return errors.New("no register redis object")
	}
}

// 依照key值,將對應的值直接設定為指定的值
func (r *RedisCliect) GetSet(dbIdx int, key string, value interface{}) error {
	if v, ok := r.redis_cliect.Load(dbIdx); ok {
		client := v.(*redis.Client)
		return client.GetSet(key, value).Err()
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
func (r *RedisCliect) StoreValue(dbIdx int, key string, value interface{}, expiration time.Duration) error {
	if v, ok := r.redis_cliect.Load(dbIdx); ok {
		client := v.(*redis.Client)
		return client.Set(key, value, expiration).Err()
	} else {
		return errors.New("no register redis object")
	}
}

// Redis `GET key` command. It returns redis.Nil error when key does not exist..
func (r *RedisCliect) LoadValue(dbIdx int, key string) (string, error) {
	if v, ok := r.redis_cliect.Load(dbIdx); ok {
		client := v.(*redis.Client)
		return client.Get(key).Result()
	} else {
		return "", errors.New("no register redis object")
	}
}

/*
確認某個key值下是否有資料存在
*/
func (r *RedisCliect) CheckValue(dbIdx int, key string) (int64, error) {
	if v, ok := r.redis_cliect.Load(dbIdx); ok {
		client := v.(*redis.Client)
		return client.Exists(key).Result()
	} else {
		return 0, errors.New("no register redis object")
	}
}

/*
刪除某一個key值
*/
func (r *RedisCliect) DeleteValue(dbIdx int, key string) error {
	if v, ok := r.redis_cliect.Load(dbIdx); ok {
		client := v.(*redis.Client)
		return client.Del(key).Err()
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
func (r *RedisCliect) StoreHValue(dbIdx int, hash string, args ...string) error {
	if v, ok := r.redis_cliect.Load(dbIdx); ok {
		client := v.(*redis.Client)
		return client.HSet(hash, args).Err()
	} else {
		return errors.New("no register redis object")
	}
}

/*
取得某個hash下的某一個field
Returns the value associated with field in the hash stored at key.
*/
func (r *RedisCliect) LoadHValue(dbIdx int, hash string, field string) (string, error) {
	if v, ok := r.redis_cliect.Load(dbIdx); ok {
		client := v.(*redis.Client)
		return client.HGet(hash, field).Result()
	} else {
		return "", errors.New("no register redis object")
	}
}

/*
取得某個hash下的全部field
Returns the value associated with field in the hash stored at key.
*/
func (r *RedisCliect) LoadHAllValue(dbIdx int, hash string) (map[string]string, error) {
	if v, ok := r.redis_cliect.Load(dbIdx); ok {
		client := v.(*redis.Client)
		return client.HGetAll(hash).Result()
	} else {
		return nil, errors.New("no register redis object")
	}
}

// func (r *RedisCliect) LoadAllHValue(dbIdx int, keys []string, field string) ([]redis.Cmder, error) {
// 	if v, ok := r.redis_cliect.Load(dbIdx); ok {
// 		client := v.(*redis.Client)
// 		pipe := client.TxPipeline()
// 		for _, key := range keys {
// 			pipe.HGet(key, field)
// 		}
// 		return pipe.Exec()
// 	} else {
// 		return nil, errors.New("no register redis object")
// 	}
// }

/*
確認某個hash下的某一個field是否存在
Returns if field is an existing field in the hash stored at key.
*/
func (r *RedisCliect) CheckHValue(dbIdx int, hash string, field string) (bool, error) {
	if v, ok := r.redis_cliect.Load(dbIdx); ok {
		client := v.(*redis.Client)
		return client.HExists(hash, field).Result()
	} else {
		return false, errors.New("no register redis object")
	}
}

/*
刪除某個hash下的某一個field
*/
func (r *RedisCliect) DeleteHValue(dbIdx int, hash string, field string) error {
	if v, ok := r.redis_cliect.Load(dbIdx); ok {
		client := v.(*redis.Client)
		return client.HDel(hash, field).Err()
	} else {
		return errors.New("no register redis object")
	}
}

func (r *RedisCliect) StoreLValue(dbIdx int, key string, values []interface{}) (int64, error) {
	if v, ok := r.redis_cliect.Load(dbIdx); ok {
		client := v.(*redis.Client)
		return client.RPush(key, values...).Result()
	} else {
		return -1, errors.New("no register redis object")
	}
}

// func (r *RedisCliect) LoadAllLALLValue(dbIdx int, keys []string) ([]redis.Cmder, error) {
// 	if v, ok := r.redis_cliect.Load(dbIdx); ok {
// 		client := v.(*redis.Client)
// 		pipe := client.TxPipeline()
// 		for _, key := range keys {
// 			pipe.LRange(key, 0, -1)
// 		}
// 		return pipe.Exec()
// 	} else {
// 		return nil, errors.New("no register redis object")
// 	}
// }

/* Redis `flushdb` command.

刪除指定redis DB內的儲存值
*/
func (r *RedisCliect) FlushDB(dbIdx int) error {
	if v, ok := r.redis_cliect.Load(dbIdx); ok {
		client := v.(*redis.Client)
		return client.FlushDB().Err()
	} else {
		return errors.New("no register redis object")
	}
}

/* Redis `flushdb` command.

刪除指定redis DB內的儲存值
*/
func (r *RedisCliect) Del(dbIdx int, key string) error {
	if v, ok := r.redis_cliect.Load(dbIdx); ok {
		client := v.(*redis.Client)
		newkey := key + "_del"
		client.Rename(key, newkey)
		return client.Del(newkey).Err()
	} else {
		return errors.New("no register redis object")
	}
}

/* Redis `Rename` command.

更改指定redis DB內的 key 值
*/
func (r *RedisCliect) Rename(dbIdx int, key, newkey string) error {
	if v, ok := r.redis_cliect.Load(dbIdx); ok {
		client := v.(*redis.Client)
		return client.Rename(key, newkey).Err()
	} else {
		return errors.New("no register redis object")
	}
}
