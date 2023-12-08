package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	_, _ = fmt.Fprintf(writer, "Hello, %s", name)
}

func GreetHttpHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "Mir")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(GreetHttpHandler)))
}
