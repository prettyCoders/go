package main

import "fmt"

func main() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	//c <- 2  //fatal error: all goroutines are asleep - deadlock!
	fmt.Println(<-c)
	fmt.Println(<-c)
}
