package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type WordCounter int
type LineCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	s := bufio.NewScanner(bytes.NewReader(p))
	s.Split(bufio.ScanWords)

	t := *c
	for s.Scan() {
		*c += 1
	}
	return int(*c - t), nil
}

func (c *LineCounter) Write(p []byte) (int, error) {
	s := bufio.NewScanner(bytes.NewReader(p))
	s.Split(bufio.ScanLines)
	t := *c
	for s.Scan() {
		*c += 1
	}
	return int(*c - t), nil
}
func main() {

	var c WordCounter
	fmt.Fprintf(&c, "name, %s", "aleksei")
	fmt.Println(c) // 13

	fmt.Fprintf(&c, "last name, %s", "makarov")
	fmt.Println(c)

	var b LineCounter
	fmt.Fprint(&b, "name: aleksei \n last_name: makarov \n")
	fmt.Println(b)

}
