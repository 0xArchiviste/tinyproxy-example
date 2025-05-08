package main

import (
	"fmt"
	"log"
	"net/http"
)

func ipcheckHandler(w http.ResponseWriter, r *http.Request) {
	clientIP := r.RemoteAddr
	fmt.Fprintf(w, "Client IP: %s\n", clientIP)
}

func main() {
	http.HandleFunc("/ipcheck", ipcheckHandler)
	
	port := ":8080"
	fmt.Printf("Server starting on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
} 