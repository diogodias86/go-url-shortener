package main

import (
	"fmt"
	"go-url-shortener/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.IndexHandler)
	//http.HandleFunc("/favicon.ico", nil)

	fmt.Println("Server running on http://localhost:8080...")
	http.ListenAndServe(":8080", nil)
}
