package core

import "math"

func Round (value float64) float64 {
	if (value < 0) {
		return math.Floor(value + .49999999999)

	} else {
		return math.Floor(value + .5)
	}
}

func MaxInt (first, second int) int {
	if (first > second) {
		return first
	} else {
		return second
	}
}
