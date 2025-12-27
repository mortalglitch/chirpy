package main

import (
	"context"
	"fmt"
	"net/http"
)

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
	if cfg.platform == "dev" {
		cfg.fileserverHits.Store(0)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hits reset to 0"))
		err := cfg.db.DeleteUsers(context.Background())
		if err != nil {
			fmt.Errorf("couldn't delete users: %w", err)
		}
		fmt.Println("Database reset successfully!")		
	}
	if cfg.platform != "dev" {
		w.WriteHeader(403)
	}

}
