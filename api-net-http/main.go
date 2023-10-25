package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	httpObj := NewHttpInterface()

	http.Handle("/foo", httpObj)

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

type HttpStructImplementation struct {
}

func NewHttpInterface() http.Handler {
	return &HttpStructImplementation{}
}

func (h *HttpStructImplementation) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("123"))
	return
}
