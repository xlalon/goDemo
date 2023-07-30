package rand

import (
	"math/rand"
	"strconv"
	"time"
)

func UIntNBetween(min, max int64) int64 {
	rand.Seed(time.Now().Unix())
	return rand.Int63n(max-min) + min
}

func DigitalMemo() string {
	return strconv.FormatInt(UIntNBetween(100000000000, 1000000000000), 10)
}
