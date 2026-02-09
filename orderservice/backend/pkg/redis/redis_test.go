package redis_test

import (
	"backend/pkg/redis"
	"testing"
)

func TestRedis(t *testing.T) {

	address := "127.0.0.1:6379"
	pwd := ""

	rc := redis.NewRedisCliectTest()

	for i := 0; i < 16; i++ {
		err := rc.CreateRedisObject(i, address, pwd)
		if err != nil {
			t.Logf("Redis database ping error: %v", err)
		}
	}

	vals := make([]interface{}, 0)

	vals = append(vals, "123")
	vals = append(vals, "456")
	vals = append(vals, "789")
	vals = append(vals, "000")

	rc.StoreLValue(2, "test", vals)

}
