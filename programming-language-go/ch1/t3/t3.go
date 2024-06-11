package main

import (
	"os"
	"strings"
)

func sepJoin() {
	var s, sep = "", ""
	for _, a := range os.Args[0:] {
		s += sep + a
		sep = " "
	}
}

func builtinJoin() {
	strings.Join(os.Args[0:], " ")
}

func main() {
	sepJoin()
	builtinJoin()
}
