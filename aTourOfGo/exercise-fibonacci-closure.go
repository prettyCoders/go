package main

import "fmt"

func fibonacci() func() int {
	pre := 0
	now := 1
	return func() int {
		now = pre + now
		pre = now - pre
		return pre
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
