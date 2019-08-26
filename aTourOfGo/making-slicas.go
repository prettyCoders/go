package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice2("a", a)

	b := make([]int, 0, 5)
	printSlice2("b", b)

	c := b[:2]
	printSlice2("c", c)

	d := c[2:5]
	printSlice2("d", d)
}

func printSlice2(name string, slice []int) {
	fmt.Println(name, slice)
	fmt.Printf("%s,len=%d,cap=%d\n", name, len(slice), cap(slice))
}
