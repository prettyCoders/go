package main

import "fmt"

type Vertex4 struct {
	X, Y int
}

func main() {
	v1 := Vertex4{1, 2}
	v2 := Vertex4{X: 3}
	v3 := Vertex4{}
	p := &Vertex4{5, 6}

	fmt.Println(v1, v2, v3, p)
}
