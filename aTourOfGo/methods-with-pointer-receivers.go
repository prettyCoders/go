package main

import (
	"fmt"
	"math"
)

type Vertex13 struct {
	X, Y float64
}

func (v *Vertex13) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex13) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

func main() {
	v := &Vertex13{
		X: 3,
		Y: 4,
	}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
