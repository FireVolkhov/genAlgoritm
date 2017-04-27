package core

import (
	"math/rand"
)

func RandomInt (start, end int) int {
	return int(Round(rand.Float64() * float64(end - start))) + start
}

func RandomBool () bool {
	return RandomInt(0, 1) == 1
}

func GetItem (array []interface{}) interface{} {
	return array[RandomInt(0, len(array) - 1)]
}
