package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func fetch(url string, ch chan<- string) {

	s := time.Now()

	r, rErr := http.Get(url)
	if rErr != nil {
		ch <- rErr.Error()
		return
	}
	defer r.Body.Close()

	n, nErr := io.Copy(io.Discard, r.Body)
	if nErr != nil {
		ch <- nErr.Error()
	}

	ch <- fmt.Sprintf("%.2fs %7d %s", time.Since(s).Seconds(), n, url)

}

func main() {

	var (
		s = time.Now()
		c = make(chan string)
	)

	for _, url := range os.Args[1:] {
		go fetch(url, c)
	}
	for range os.Args[1:] {
		fmt.Println(<-c)
	}

	fmt.Printf("%.2fs elapsed \n", time.Since(s).Seconds())

}
