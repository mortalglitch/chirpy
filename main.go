package main

import (
	"sync/atomic"
	"fmt"
	"io"
	"log"
	"net/http"
)

type apiConfig struct {
	fileserverHits atomic.Int32
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(w, r)
	})
}

func (cfg *apiConfig) handlerMetrics(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	amount := cfg.fileserverHits.Load()
	io.WriteString(w, fmt.Sprintf("Hits: %d\n", amount))
}

func (cfg *apiConfig) resetCount(w http.ResponseWriter, _ *http.Request){
	cfg.fileserverHits.Store(0)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	fmt.Println("Starting Server...")
	mux := http.NewServeMux()
	apiCfg := apiConfig{}
	healthHandler := func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}
	mux.HandleFunc("/healthz", healthHandler)
	mux.HandleFunc("/metrics", apiCfg.handlerMetrics)
	mux.HandleFunc("/reset", apiCfg.resetCount)
	//mux.Handle("/app/", http.StripPrefix("/app/", http.FileServer(http.Dir("."))))
	mux.Handle("/app/", apiCfg.middlewareMetricsInc(http.StripPrefix("/app/", http.FileServer(http.Dir(".")))))

	server := &http.Server {
		Handler:  mux,
		Addr:     ":8080",
	}
	
	log.Printf("Serving on post: %s\n", server.Addr)
	log.Fatal(server.ListenAndServe())

}



