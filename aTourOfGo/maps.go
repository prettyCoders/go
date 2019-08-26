package main

import "fmt"

type Vertex5 struct {
	lat, long float64
}

var m map[string]Vertex5

func main() {
	m = make(map[string]Vertex5)
	m["Bell Labs"] = Vertex5{
		lat:  40.68433,
		long: -74.39967,
	}

	fmt.Println(m["Bell Labs"])
}
