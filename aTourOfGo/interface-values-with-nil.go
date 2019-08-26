package main

import "fmt"

type I3 interface {
	M()
}

type T3 struct {
	s string
}

func (t *T3) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.s)
}

func describe2(i I3) {
	fmt.Printf("(%v,%T)\n", i, i)
}

func main() {
	var t *T3
	var i I3 = t
	describe2(i)
	i.M()

	i = &T3{"hello"}
	describe2(i)
	i.M()
}
