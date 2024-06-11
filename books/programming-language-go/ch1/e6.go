package main

import (
	"bufio"
	"fmt"
	"os"
)

func countLines(f *os.File, c map[string]int) {
	i := bufio.NewScanner(f)
	for i.Scan() {
		c[i.Text()]++
	}
}

func main() {
	c := make(map[string]int)
	f := os.Args[1:]
	if len(f) == 0 {
		countLines(os.Stdin, c)
	} else {
		for _, n := range f {
			o, oErr := os.Open(n)
			if oErr != nil {
				fmt.Fprintf(
					os.Stderr,
					"error:%s\n",
					oErr.Error(),
				)
				continue
			}
			countLines(o, c)
			o.Close()
		}
	}
	for l, n := range c {
		if n > 1 {
			fmt.Printf("%s -- %d \n", l, n)
		}
	}
}
