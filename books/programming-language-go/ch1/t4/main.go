package main

import (
	"fmt"
	"os"
	"strings"
)

func exists(a []string, k string) bool {
	for _, v := range a {
		if k == v {
			return true
		}
	}
	return false
}

func main() {

	c := make(map[string]int)
	a := make(map[string][]string)

	for _, n := range os.Args[1:] {

		f, fErr := os.ReadFile(n)
		if fErr != nil {
			panic(fErr)
		}

		for _, l := range strings.Split(string(f), "\n") {
			c[l]++
			_, e := a[l]
			if !e {
				a[l] = append(a[l], n)
			} else {
				if !exists(a[l], n) {
					a[l] = append(a[l], n)
				}
			}
		}

	}

	for l, n := range c {
		if n > 1 {
			fmt.Printf("file(s): %s; line: %s; total: %d \n", strings.Join(a[l], ", "), l, n)
		}
	}

}
