package main

import (
	"fmt"
	"net/http"
	)

func main() {
	http.HandleFunc("/", func(w  http.ResponseWriter,   r *http.Request) {
		fmt.Fprintf(w, "Hello from my app running on Kubernetes!")
	})
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Status: Healthy")
	})
	fmt.Println("Server starting on port 8080...")
	http.ListenAndServe(":8080", nil)
}
