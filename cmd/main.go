package main

import (
	"log"
	"net/http"

	"hello-api/handlers/rest"
)

func main() {
	addr := ":8081"

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", rest.TranslateHandler)

	log.Printf("listening on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
