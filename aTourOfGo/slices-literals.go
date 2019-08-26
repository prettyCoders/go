package main

import "fmt"

func main() {
	q := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(q)
	r := []bool{true, false, true, false, true, false}
	fmt.Println(r)

	s := []struct {
		i int
		j bool
	}{
		{1, true},
		{2, false},
		{3, true},
		{4, false},
		{5, true},
		{6, false},
	}
	fmt.Println(s)
}
