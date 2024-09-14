package main

import (
	"bufio"
	"fmt"
	"os"
)

func solution(arr []int) int {
	if len(arr) == 1 {
		return 1
	}

	var num = 0

	for i := 0; i < len(arr)-1; i++ {
		var tpNum = 1
		var known = map[int]bool{
			arr[i]: true,
		}
		for j := i + 1; j < len(arr); j++ {
			known[arr[j]] = true
			if len(known) <= 2 {
				tpNum++
			} else {
				break
			}
		}
		if tpNum > num {
			num = tpNum
		}
	}

	return num
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

		var arr []int

		for j := 0; j < n; j++ {
			var aj int
			fmt.Fscan(ri, &aj)
			arr = append(arr, aj)
		}

		// here task solution starts
		//
		// child -> parent
		fmt.Fprintf(wo, "%d\n", solution(arr))

	}
}
