package main

import (
	"context"
	"net/http"
	"time"

	"github.com/mortalglitch/chirpy/internal/auth"
	"github.com/mortalglitch/chirpy/internal/database"
)

func (cfg *apiConfig) handlerRefreshCheck(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Token string `json:"token"`
	}

	checkRefreshToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't fetch token: ", err)
		return
	}

	checkUser, err := cfg.db.GetUserFromRefreshToken(context.Background(), checkRefreshToken)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Error with refresh token: ", err)
		return
	}

	if checkUser.ExpiresAt.Before(time.Now().UTC()) || checkUser.RevokedAt.Valid {
		respondWithError(w, http.StatusUnauthorized, "Error with refresh token: ", err)
		return
	}

	newJWT, err := auth.MakeJWT(checkUser.UserID, cfg.signingKey, time.Duration(3600)*time.Second)
	if err != nil {
		respondWithError(w, 400, "Unable to generate JWT for current user: ", err)
	}

	newToken := response{
		Token: newJWT,
	}
	respondWithJSON(w, 200, newToken)
}

func (cfg *apiConfig) handlerRevokeRefresh(w http.ResponseWriter, r *http.Request) {
	checkRefreshToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't fetch token: ", err)
		return
	}

	confirmRevoked := cfg.db.MarkRefreshRevoked(context.Background(), database.MarkRefreshRevokedParams{
		UpdatedAt: time.Now().UTC(),
		Token:     checkRefreshToken,
	})
	if confirmRevoked != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't revoke token: ", err)
		return
	}

	w.WriteHeader(204)
}
