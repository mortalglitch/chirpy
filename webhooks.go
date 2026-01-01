package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/mortalglitch/chirpy/internal/auth"
	"github.com/mortalglitch/chirpy/internal/database"
)

type PolkaRedUpgrade struct {
	Event string        `json:"event"`
	Data  UserReference `json:"data"`
}

type UserReference struct {
	UserID uuid.UUID `json:"user_id"`
}

func (cfg *apiConfig) handlerRedUpgradeUser(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	params := PolkaRedUpgrade{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	requestKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, 401, "Unable to fetch API Key from request", err)
		return
	}

	if requestKey != cfg.polkakey {
		respondWithError(w, 401, "Authorization Error", err)
	}

	if params.Event != "user.upgraded" {
		w.WriteHeader(204)
		return
	}

	userUpgrade := cfg.db.UpgradeUserToRed(context.Background(), database.UpgradeUserToRedParams{
		UpdatedAt: time.Now().UTC(),
		ID:        params.Data.UserID,
	})
	if userUpgrade != nil {
		w.WriteHeader(404)
		return
	}

	w.WriteHeader(204)
}
