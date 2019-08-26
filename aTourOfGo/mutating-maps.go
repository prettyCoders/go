package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	value, ok := m["Answer"]
	fmt.Println("The value", value, "present?", ok)
}
