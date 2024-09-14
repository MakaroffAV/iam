package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findMax(numbers []int) int {
	m := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if numbers[i] > m {
			m = numbers[i]
		}
	}
	return m
}

func main() {
	ri := bufio.NewReader(os.Stdin)
	wo := bufio.NewWriter(os.Stdout)
	defer wo.Flush()

	var t int
	fmt.Fscan(ri, &t)

	for i := 0; i < t; i++ {
		var n int64
		fmt.Fscan(ri, &n)

		number := fmt.Sprintf("%d", n)
		fmt.Println(number)
		if len(number) == 1 {
			fmt.Fprintf(wo, "%d\n", 0)
		} else {
			var numbers []int
			split_number := strings.Split(number, "")
			for i := 0; i < len(split_number); i++ {
				var temp_numbers []string
				for j := 0; j < len(split_number); j++ {
					if i != j {
						temp_numbers = append(temp_numbers, split_number[j])
					}
				}
				number_int, err := strconv.Atoi(strings.Join(temp_numbers, ""))
				if err != nil {
					panic(err)
				}
				numbers = append(numbers, number_int)
			}
			fmt.Fprintf(wo, "%d\n", findMax(numbers))
		}
	}
}
