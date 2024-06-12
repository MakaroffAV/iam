package main

import (
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

		_, cErr := io.Copy(os.Stdout, r.Body)
		if cErr != nil {
			panic(cErr)
		}

	}
}
