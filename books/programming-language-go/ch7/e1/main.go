package main

import "fmt"

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}
func main() {

	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // 5

	c = 0

	// w   io.Writer   интерфейс,   который
	// гарантирует реализацию метода  Write
	//
	// Мы передаем свой тип, в котором есть
	// реализация        метода       Write
	fmt.Fprintf(&c, "name, %s", "aleksei")
	fmt.Println(c) // 13

}
