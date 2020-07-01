package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/diogodias86/go-url-shortener/handlers"
)

func main() {
	fs := http.FileServer(http.Dir("./templates"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/favicon.ico", handlers.FavIconHandler)
	http.HandleFunc("/", handlers.IndexHandler)

	fmt.Println("Server running on http://localhost:8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
