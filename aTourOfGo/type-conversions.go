package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y = 3, 4
	var f = math.Sqrt(float64(x*x + y*y))
	var z = uint(f)
	fmt.Println(x, y, z)
}
