package main

import (
	"log"
	"net/http"

	"practice/internal/application"
)

func main() {
	mux := application.New()
	log.Fatal(http.ListenAndServe(":8080", mux))
}
