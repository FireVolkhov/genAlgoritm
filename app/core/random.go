package core

import (
	"math/rand"
	"math"
	"time"
)

func RandomInt (start, end int) int {
	diff := end - start
	return int(Round(rand.Float64() * float64(diff))) + start
}

func RandomBool () bool {
	return RandomInt(0, 1) == 1
}

func RandomFloat64 () float64 {
	return rand.Float64()
}

func GetItem (array []interface{}) interface{} {
	return array[RandomInt(0, len(array) - 1)]
}

//func init () {
//	go setRandomSeed()
//}

func setRandomSeed() {
	for {
		rand.Seed(int64(rand.Float64() * math.MaxInt64))
		time.Sleep(time.Millisecond * 100)
	}
}