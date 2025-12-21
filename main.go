package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting Server...")
	mux := http.NewServeMux()

	server := &http.Server {
		Handler:  mux,
		Addr:     ":8080",
	}
	
	log.Printf("Serving on post: %s\n", server.Addr)
	log.Fatal(server.ListenAndServe())

}

