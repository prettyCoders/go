package main

import (
	"fmt"
	"math"
)

var temps []float64

func isOldValue(value float64) (isOld bool) {
	for _, v := range temps {
		if value == v {
			return true
		}
	}
	return false
}

func sqrtTest(x float64) float64 {
	z := float64(1)
	for i := 0; i < 10; i++ {
		temp := (z*z - x) / (2 * z)
		if isOldValue(temp) {
			return z
		}
		temps = append(temps, temp)
		z -= temp
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(sqrtTest(895))
	fmt.Println(math.Sqrt(895))
}
