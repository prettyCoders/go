package main

import "fmt"

func main() {
	var i interface{}
	describe4(i)

	i = 45
	describe4(i)

	i = "hello"
	describe4(i)
}

func describe4(i interface{}) {
	fmt.Printf("(%v,%T)\n", i, i)
	fmt.Print()
}
