package main

import (
	"fmt"
	"math"
)

type Vertex12 struct {
	X, Y float64
}

func (v Vertex12) abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbdFunc(v Vertex12) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex12{
		X: 3,
		Y: 4,
	}

	fmt.Println(v.abs())
	fmt.Println(AbdFunc(v))

	p := &v

	fmt.Println(p.abs())
	fmt.Println(AbdFunc(*p))
}
