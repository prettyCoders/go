package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	s := strings.NewReader("Hello Reader!")
	b := make([]byte, 8)
	for {
		n, e := s.Read(b)
		fmt.Printf("n=%v,err=%v,b=%v\n", n, e, b)
		fmt.Printf("b[:n]=%v\n", b[:n])
		if e == io.EOF {
			break
		}
	}
}
