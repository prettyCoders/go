package main

import "fmt"

type I interface {
	m()
}

type T struct {
	s string
}

func main() {
	var i I = &T{"hello"}
	i.m()
}

func (t *T) m() {
	fmt.Println(t.s)
}
