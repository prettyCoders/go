package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{
		Name: "sunlight",
		Age:  25,
	}
	z := Person{
		Name: "poul",
		Age:  42,
	}
	fmt.Println(a, z)
}
