package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Backend działa!")
	})

	log.Printf("Serwer startuje na porcie %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}