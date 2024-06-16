package main

import "fmt"

func nonEmpty(a []string) []string {
	i := 0
	for _, s := range a {
		if s != "" {
			a[i] = s
			i++
		}
	}
	return a[:i]
}

func nonEmpty2(a []string) []string {
	out := a[:0] // срез нулевой длины
	for _, s := range a {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func removeFromSlice(s []int, i int) []int {
	// s = []int{5, 6, 7, 8, 9}
	// i = 2

	// s[i:] ~ [....., 7, 8, 9]
	// s[i+1:] ~ [......, 8, 9]
	copy(s[i:], s[i+1:]) // s = [5 6 8 9 9]
	return s[:len(s)-1]
}

func main() {

	// изменяем массив  и делаем
	// новый срез, который будет
	// удоволетворять    условию
	a := []string{"one", "", "two"}
	fmt.Printf("%q\n", nonEmpty(a)) // ["one" "two"]

	// то же самое, но  с append
	b := []string{"one", "", "twoX"}
	fmt.Printf("%q %q\n", nonEmpty2(b), b) // ["one" "twoX"] ["one" "twoX" "twoX"]

	// удаление из середины стека
	c := []int{5, 6, 7, 8, 9}
	fmt.Println(removeFromSlice(c, 2))
}
