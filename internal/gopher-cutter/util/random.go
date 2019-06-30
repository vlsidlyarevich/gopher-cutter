package util

import (
	"math/rand"
	"strconv"
	"time"
)

type Random struct {

}

func NewRandom() (r *Random) {
	rand.Seed(time.Now().UnixNano())

	return new(Random)
}

func (r *Random) RInt(length int) (res int, e error) {
	var s string
	for i := 0; i < length; i++ {
		s += (string)(rand.Intn(10) + 48)
	}

	return strconv.Atoi(s)
}
