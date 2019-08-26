package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6}

	s = s[1:4]
	fmt.Println(s) //1 2 3

	s = s[:2]
	fmt.Println(s) //1 2

	s = s[1:]
	fmt.Println(s) //2
}
