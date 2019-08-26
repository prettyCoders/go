package main

import (
	"fmt"
	"math"
)

type Vertex8 struct {
	X, Y float64
}

func abs(v Vertex8) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func main() {
	v := Vertex8{
		X: 3,
		Y: 4,
	}
	fmt.Println(abs(v))
}
