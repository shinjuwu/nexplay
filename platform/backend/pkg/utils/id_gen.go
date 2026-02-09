package utils

import (
	"math"

	"go.uber.org/atomic"
)

/*
流水號 ID 產生器
*/
type SerialNumberGen struct {
	initValue     int32
	serialCounter *atomic.Int32
}

func NewSerialNumberGen(init int32) *SerialNumberGen {
	return &SerialNumberGen{
		initValue:     init,
		serialCounter: atomic.NewInt32(init),
	}
}

func (b *SerialNumberGen) IncrCounter() int {
	if b.serialCounter.Load() == math.MaxInt32 {
		b.serialCounter.Store(b.initValue)
	}
	b.serialCounter.Inc()
	return int(b.serialCounter.Load())
}
