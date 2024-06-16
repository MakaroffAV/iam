package main

import (
	"fmt"
	"time"
)

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func spinner(d time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(d)
		}
	}
}

func main() {

	// вот  эта подпрограмма
	// живет пока живет main
	go spinner(100 * time.Millisecond)

	// cчитаем   45-е  число
	const n = 45
	fibN := fib(45)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}
