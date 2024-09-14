package main

import (
	"bufio"
	"fmt"
	"os"
)

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
		var indx = 0
		var tree = map[int]int{}

		for indx < n-1 {

			if arr[indx+1] == 0 {
				if tree[arr[indx]] == 0 {
					tree[arr[indx]] = 0
				}
				indx += 2
			} else {
				root := arr[indx]
				inum := arr[indx+1]
				for i1 := 0; i1 < inum; i1++ {
					// fmt.Println("i am here", root, arr[indx+1+1+i1], indx+1+1+i1)
					tree[arr[indx+1+1+i1]] = root
				}
				indx = indx + 1 + inum + 1
			}
		}

		treeRoot := 0
		if len(tree) == 1 {
			for k, v := range tree {
				if v != 0 {
					fmt.Fprintf(wo, "%d\n", v)
				} else {
					fmt.Fprintf(wo, "%d\n", k)
				}
			}
		} else {
			for _, v := range tree {
				_, ok := tree[v]
				if !ok {
					treeRoot = v
				}
			}
			fmt.Fprintf(wo, "%d\n", treeRoot)
		}
	}
}
