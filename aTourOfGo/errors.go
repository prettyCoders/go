package main

import (
	"fmt"
	"strconv"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (err *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", err.When, err.What)
}

func run() error {
	return &MyError{When: time.Now(), What: "It did`t work"}
}

func main() {
	i, e := strconv.Atoi("45")
	fmt.Println(i, e)

	if err := run(); err != nil {
		fmt.Println(err)
	}
}
