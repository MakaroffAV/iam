package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	c := make(map[string]int)
	i := bufio.NewScanner(os.Stdin)
	for i.Scan() {
		c[i.Text()]++
	}
	for l, n := range c {
		if n > 1 {
			fmt.Printf("%s - %d \n", l, n)
		}
	}
}
