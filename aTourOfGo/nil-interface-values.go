package main

import "fmt"

type I4 interface {
	M()
}

func describe3(i I4) {
	fmt.Printf("(%v,%T)\n", i, i)
}

func main() {
	var i I4
	describe3(i)
	i.M()
}
