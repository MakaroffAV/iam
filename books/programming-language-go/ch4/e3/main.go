package main

import "fmt"

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// Имеется место для роста, расширяем срез
		z = x[:zlen]
	} else {
		// Места для роста нет. Выделяем новый массив,
		// увеличиваем его вместимость в два раза
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func main() {

	var runes []rune
	for _, r := range "hello, 梅里" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)

	r := []rune("hello, 梅里")
	fmt.Printf("%q\n", r)

	a := make([]int, 2, 4)
	fmt.Println(a) // [0, 0]

	a = append(a, 1)
	a = append(a, 2)
	fmt.Println(a) // [0, 0, 1, 2]

	a = append(a, 3)
	a = append(a, 4)
	fmt.Println(a) // [0, 0, 1, 2, 3]

	// Своя функция для расширения
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}

	fmt.Println()

	// Встроенная функция для расширения
	var xx = make([]int, 0, 1)
	for i := 0; i < 10; i++ {
		xx = append(xx, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(xx), xx)
	}

}
