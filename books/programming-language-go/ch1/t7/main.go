package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {

		var u string
		if strings.HasPrefix(url, "https://") {
			u = url
		} else {
			u = "https://" + url
		}

		r, rErr := http.Get(u)
		if rErr != nil {
			panic(rErr)
		}
		defer r.Body.Close()

		fmt.Printf("%s - %d \n", u, r.StatusCode)

	}
}
