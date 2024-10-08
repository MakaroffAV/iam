package main

import "fmt"

func main() {
	n := make(chan int)
	s := make(chan int)
	go func() {
		for x := 0; x < 100; x++ {
			n <- x
		}
		close(n)
	}()
	go func() {
		for x := range n {
			s <- x * x
		}
		close(s)

	}()

	for x := range s {
		fmt.Println(x)
	}

}
