package main

import (
	"crypto/sha256"
	"fmt"
)

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RUR
)

func main() {

	var a [3]int

	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	for i, v := range a {
		fmt.Printf("%d %d \n", i, v)
	}

	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(q)
	fmt.Println(r)

	x := [...]int{1, 1, 1}
	fmt.Println(x)

	symbol := [...]string{
		USD: "$",
		EUR: "€",
		GBP: "£",
		RUR: "₽",
	}
	fmt.Println(RUR, symbol[RUR])

	a1 := [2]int{1, 2}
	b1 := [...]int{2, 1}
	fmt.Println(a1 == b1) // false

	a2 := [2]int{1, 2}
	b2 := [...]int{1, 2}
	fmt.Println(a2 == b2) // true

	h1 := sha256.Sum256([]byte("x"))
	h2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%t\n%T\n%T\n%x\n%x\n", h1 == h2, h1, h2, h1, h2) // false, [32]uint8, [32]uint8

	j1 := [...]int{1, 2}
	j1[0] = 10
	fmt.Println(j1)

}
