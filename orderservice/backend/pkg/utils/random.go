package utils

import (
	"backend/pkg/encrypt/md5hash"
	"fmt"
	"math"
	"math/rand"
	"time"
)

// RandomIntMinToMax random range is 0~N-1
func RandomIntMinToMax(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

// Randomly generate a random number string of specified length.
func RomdomString(length int) string {
	num := ToInt32(math.Pow10(length+1), 0)
	format := fmt.Sprintf("%%0"+"%d"+"v", 4)
	return fmt.Sprintf(format, rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(num))
}

// create recharge order id.
func CreateOrderId() string {
	return GetUnsignedTimeNowUTC() + RomdomString(4)
}

func CreatreOrderIdByOrderTypeAndSalt(orderType int, salt string, betTime time.Time) string {
	str := fmt.Sprintf("%d%d%02d%02d%02d%02d%02d%06d", orderType, betTime.Year(), betTime.Month(), betTime.Day(), betTime.Hour(),
		betTime.Minute(), betTime.Second(), betTime.UnixMilli()%int64(time.Millisecond))
	hash := md5hash.Hash32bit(salt)

	c := 0
	for i := 0; i < len(hash); i++ {
		if hash[i] >= '0' && hash[i] <= '9' {
			c++
			str += string(hash[i])
			if c == 6 {
				break
			}
		}
	}

	for ; c < 6; c++ {
		str += "0"
	}

	return str
}
