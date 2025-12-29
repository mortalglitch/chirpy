package main

import (
	"context"
	"time"
	"encoding/json"
	"net/http"

	"github.com/mortalglitch/chirpy/internal/database"
	"github.com/mortalglitch/chirpy/internal/auth"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
}


func (cfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	hashedPassword, err := auth.HashPassword(params.Password)
	if err != nil{
		respondWithError(w, http.StatusInternalServerError, "Couldn't hash password", err)
		return
	}
	
	user, err := cfg.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:              uuid.New(),
		CreatedAt:       time.Now().UTC(),
		UpdatedAt:       time.Now().UTC(),
		Email:           params.Email,
		HashedPassword:  hashedPassword,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create user", err)
		return
	}

	newUser := User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
	}
	
	respondWithJSON(w, 201, newUser) 
}

func (cfg *apiConfig) handlerUserLogin(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	userData, err := cfg.db.GetUserByEmail(context.Background(), params.Email)
	if err != nil {
		respondWithError(w, 401, "Incorrect email or password", err)
		return
	}

	validPass, err := auth.CheckPasswordHash(params.Password, userData.HashedPassword)
	if err != nil {
		respondWithError(w, 401, "Incorrect email or password", err)
		return
	}

	if !validPass {
		respondWithError(w, 401, "Incorrect email or password", err)
		return
	}

	if validPass {
		loggedUser := User{
			ID:        userData.ID,
			CreatedAt: userData.CreatedAt,
			UpdatedAt: userData.UpdatedAt,
			Email:     userData.Email,
		}
	
		respondWithJSON(w, 200, loggedUser)
	}

}
