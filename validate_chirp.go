package main

import (
	"encoding/json"
	"net/http"
	"slices"
	"strings"
)

func handlerValidateChirp(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}
	type returnVals struct {
		CleanedBody string `json:"cleaned_body"`
	}


	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	const maxChirpLength = 140
	if len(params.Body) > maxChirpLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}

	cleanBody := cleanChirp(params.Body)

	respondWithJSON(w, http.StatusOK, returnVals{
		CleanedBody: cleanBody,
	})
}

func cleanChirp(body string) string {
	profane := []string{"kerfuffle", "sharbert", "fornax"}

	splitBody := strings.Split(body, " ")
	for i, item := range splitBody {
		found := slices.Contains(profane, strings.ToLower(item))
		if found {
			splitBody[i] = "****"
		}
	}

	return strings.Join(splitBody, " ")
}
