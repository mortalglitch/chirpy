package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting Server...")
	mux := http.NewServeMux()
	
	healthHandler := func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}
	mux.HandleFunc("/healthz", healthHandler)

	mux.Handle("/app/", http.StripPrefix("/app/", http.FileServer(http.Dir("."))))

	server := &http.Server {
		Handler:  mux,
		Addr:     ":8080",
	}
	
	log.Printf("Serving on post: %s\n", server.Addr)
	log.Fatal(server.ListenAndServe())

}


