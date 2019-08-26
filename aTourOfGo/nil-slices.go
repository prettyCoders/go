package main

import "fmt"

func main() {
	var s []int
	fmt.Printf("len=%d,cap=%d\n", len(s), cap(s))
	if s == nil {
		fmt.Println("nil")
	}
}
