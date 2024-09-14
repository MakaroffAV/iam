package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func main() {
	ri := bufio.NewReader(os.Stdin)
	wo := bufio.NewWriter(os.Stdout)
	defer wo.Flush()

	var t int
	fmt.Fscan(ri, &t)

	for i := 0; i < t; i++ {
		var n int
		var p float64
		var miss float64
		fmt.Fscan(ri, &n, &p)
		for j := 0; j < n; j++ {
			var price float64
			fmt.Fscan(ri, &price)
			miss += math.Mod(math.Round((price*(p/100))*100)/100, 1)
		}
		fmt.Fprintf(wo, "%.2f\n", miss)
	}
}
