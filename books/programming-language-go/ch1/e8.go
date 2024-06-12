package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		r, rErr := http.Get(url)
		if rErr != nil {
			panic(rErr)
		}
		defer r.Body.Close()

		b, bErr := io.ReadAll(r.Body)
		if bErr != nil {
			panic(bErr)
		}

		fmt.Printf("%s", b)
	}
}
