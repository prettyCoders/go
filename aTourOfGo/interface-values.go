package main

import "fmt"

type I2 interface {
	m2()
}

type T2 struct {
	S string
}

func (t *T2) m2() {
	fmt.Println(t.S)
}

type F float64

func (f F) m2() {
	fmt.Println(f)
}

func describe(i I2) {
	fmt.Printf("(%v,%T)\n", i, i)
}

func main() {
	var i I2
	i = &T2{"hello"}
	describe(i)
	i.m2()

	i = F(0.3)
	describe(i)
	i.m2()
}
