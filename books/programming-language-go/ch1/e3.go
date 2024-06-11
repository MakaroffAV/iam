package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, a := range os.Args[1:] {
		s += sep + a
		sep = " "
	}
	fmt.Println(s)
}
