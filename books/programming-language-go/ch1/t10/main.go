package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const url = "https://zakupki.gov.ru/epz/order/notice/ea20/view/common-info.html?regNumber=0322300002524000051"

func main() {

	f, fErr := os.Create("output")
	if fErr != nil {
		panic(fErr)
	}
	defer f.Close()

	t := time.Now()
	r, rErr := http.Get(url)
	if rErr != nil {
		panic(rErr)
	}
	defer r.Body.Close()

	_, bErr := io.Copy(f, r.Body)
	if bErr != nil {
		panic(bErr)
	}

	fmt.Printf("%.2fs elapsed", time.Since(t).Seconds())

}
