package main

import (
	"fmt"
	"math"
)

type Vertex9 struct {
	X, Y float64
}

func (v Vertex9) abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex9) scale(f float64) {
	v.X *= f
	v.Y *= f
}

func main() {
	v := Vertex9{
		X: 3,
		Y: 4,
	}
	v.scale(10)
	fmt.Println(v.abs())
}
