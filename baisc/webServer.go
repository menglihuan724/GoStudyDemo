package main

import (
	"fmt"
	"log"
	"net/http"
)

type Hello struct{}
type Hello2 struct{}
type String string

func (h Hello) ServeHTTP(
w http.ResponseWriter,
r *http.Request) {
	fmt.Fprint(w, String("123"))
}

func (h Hello2) ServeHTTP(
w http.ResponseWriter,
r *http.Request) {
	fmt.Fprint(w, String("456"))
}

func main() {

	var h Hello
	var h2 Hello2
	http.Handle("/string", h)
	http.Handle("/test", h2)
	err := http.ListenAndServe("localhost:4000", h)
	if err != nil {
		log.Fatal(err)
	}
}
