package main

import (
	"fmt"
	"math"
)

type Vertex10 struct {
	X, Y float64
}

func vbs(v Vertex10) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func scale(v *Vertex10, f float64) {
	v.X *= f
	v.Y *= f
}

func main() {
	v := Vertex10{
		X: 3,
		Y: 4,
	}
	scale(&v, 10)
	fmt.Println(vbs(v))
}
