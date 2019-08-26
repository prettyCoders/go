package main

import "fmt"

type Vertex11 struct {
	X, Y float64
}

func (v *Vertex11) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

func FuncScale(v *Vertex11, f float64) {
	v.X *= f
	v.Y *= f
}

func main() {
	v := Vertex11{
		X: 3,
		Y: 4,
	}
	v.Scale(2)
	FuncScale(&v, 10)

	p := &Vertex11{
		X: 4,
		Y: 3,
	}
	p.Scale(3)
	FuncScale(p, 8)

	fmt.Println(v, p)

}
