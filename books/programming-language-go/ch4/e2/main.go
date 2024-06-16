package main

import (
	"fmt"
)

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {

	a := make([]int, 10, 20)
	fmt.Println(a[:15])
	// fmt.Println(a[:25]) // panic

	b := [10]int{2, 3, 4}
	c := b[4:]
	c[0] = 100
	fmt.Println(b)

	d := [...]int{1, 2, 3, 4}
	reverse(d[:])
	fmt.Println(d)

	// циклический    сдвиг
	// влево на две позиции
	//
	// [0, 1, 2, 3, 4, 5]
	// [1, 0, 2, 3, 4, 5]
	// [1, 0, 5, 4, 3, 2]
	//
	// [2, 3, 4, 5, 0, 1]
	e := []int{0, 1, 2, 3, 4, 5}
	reverse(e[:2])
	reverse(e[2:])
	reverse(e)
	fmt.Println(e)

	// f := []int{1, 2, 3}
	// g := []int{1, 2, 3}
	// fmt.Println(f == g) // compilation error

}
