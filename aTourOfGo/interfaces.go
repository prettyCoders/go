package main

import (
	"fmt"
	"math"
)

type MyFloat2 float64

type Vertex14 struct {
	X, Y float64
}

type Abser interface {
	abs() float64
}

func (v *Vertex14) abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (f MyFloat2) abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	var a Abser
	v := Vertex14{3, 4}
	f := MyFloat2(-math.Sqrt2)

	a = f
	a = &v

	fmt.Println(a.abs())
}
