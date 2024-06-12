package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%s]=%s\n", k, v)
	}
	fmt.Fprintf(w, "Host=%s\n", r.Host)
	fmt.Fprintf(w, "RemoteAddress=%s\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%s]=%s\n", k, v)
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
