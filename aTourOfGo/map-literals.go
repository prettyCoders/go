package main

import "fmt"

type Vertex6 struct {
	lab, long float64
}

var m2 = map[string]Vertex6{
	"Bell Labs": {
		lab:  1.5,
		long: -9.8,
	},
	"Google": {
		lab:  758.1,
		long: -3545.4,
	},
}

func main() {
	fmt.Println(m2)
}
