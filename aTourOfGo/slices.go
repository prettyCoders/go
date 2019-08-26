package main

import "fmt"

func main() {
	primes := [6]int{2, 8, 5, 64, 47, 1}
	var s []int = primes[1:4]
	fmt.Println(s)
}
