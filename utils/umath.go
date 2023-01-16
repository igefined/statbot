package utils

import "math"

func RoundFloat2Decimal(num float64) float64 {
	return math.Floor(num*100) / 100
}
