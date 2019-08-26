package main

import (
	"fmt"
	"math"
)

type Vertex7 struct {
	X, Y float64
}

func (v Vertex7) abs() float64 {
	return math.Sqrt(v.X*v.Y + v.Y*v.Y)
}

func main() {
	v := Vertex7{
		X: 15,
		Y: 8,
	}
	fmt.Println(v.abs())
}
