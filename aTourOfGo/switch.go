package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Go runs on!")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("darwin")
	case "linux":
		fmt.Println("linux")
	case "windows":
		fmt.Println("windows")
	default:
		fmt.Println(os)
	}
}
