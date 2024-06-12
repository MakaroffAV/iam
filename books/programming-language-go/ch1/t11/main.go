package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func fetch(url string, ch chan<- string) {

	t := time.Now()

	r, rErr := http.Get(url)
	if rErr != nil {
		ch <- rErr.Error()
	}
	defer r.Body.Close()

	_, nErr := io.Copy(io.Discard, r.Body)
	if nErr != nil {
		ch <- nErr.Error()
	}

	ch <- fmt.Sprintf("%.2fs elapsed for %s", time.Since(t).Seconds(), url)

}

func main() {

	s := time.Now()
	c := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, c)
	}

	for range os.Args[1:] {
		fmt.Println(<-c)
	}

	fmt.Printf("%.2fs all time elapsed \n", time.Since(s).Seconds())

}
