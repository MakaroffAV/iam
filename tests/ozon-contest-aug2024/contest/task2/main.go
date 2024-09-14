package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD = 1000000007

func countWays(n int, l []int, r []int) int {
	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		start := (l[i-1] + i - 1) / i
		end := r[i-1] / i
		count := max(0, end-start+1)

		if count > 0 {
			newDp := make([]int, n+1)
			for j := 0; j <= n; j++ {
				if dp[j] > 0 {
					newJ := j + 1
					if newJ <= n {
						newDp[newJ] = (newDp[newJ] + dp[j]*count) % MOD
					}
				}
			}
			dp = newDp
		}
	}

	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	ri := bufio.NewReader(os.Stdin)
	wo := bufio.NewWriter(os.Stdout)
	defer wo.Flush()

	var t int
	fmt.Fscan(ri, &t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(ri, &n)

		l := make([]int, 0, n)
		for j := 0; j < n; j++ {
			var lj int
			fmt.Fscan(ri, &lj)
			l = append(l, lj)
		}

		r := make([]int, 0, n)
		for j := 0; j < n; j++ {
			var rj int
			fmt.Fscan(ri, &rj)
			r = append(r, rj)
		}

		fmt.Fprintf(wo, "%d\n", countWays(n, l, r))
	}

	// Пример использования
	//n := 5
	//l := []int{1, 2, 7, 10, 20}
	//r := []int{1, 4, 10, 30, 40}
	//fmt.Println(countWays(n, l, r)) // Ожидаемый вывод: 4
}
