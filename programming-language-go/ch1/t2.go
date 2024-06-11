package main

import (
	"os"
	"fmt"
)

func main() {
	for i, a := range os.Args {
	fmt.Println(i, a)
	}
}
