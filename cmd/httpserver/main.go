package main

import (
	"log"
	"net/http"

	"github.com/marvin-min/testcontainers-go-demo/go_specs_greet"
)

func main() {
	handler := http.HandlerFunc(go_specs_greet.Handler)
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
