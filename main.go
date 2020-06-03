package main

import (
	"fmt"
	"go-url-shortener/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/favicon.ico", handlers.FavIconHandler)
	http.HandleFunc("/", handlers.IndexHandler)

	fmt.Println("Server running on http://localhost:8080...")
	http.ListenAndServe(":8080", nil)
}
