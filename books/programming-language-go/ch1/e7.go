package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	c := make(map[string]int)
	for _, n := range os.Args[1:] {

		d, dErr := os.ReadFile(n)
		if dErr != nil {
			fmt.Fprintf(
				os.Stderr,
				"error:%s\n",
				dErr.Error(),
			)
			continue
		}

		for _, l := range strings.Split(string(d), "\n") {
			c[l]++
		}

	}

	for l, n := range c {
		if n > 1 {
			fmt.Printf("%s -- %d \n", l, n)
		}
	}
}
