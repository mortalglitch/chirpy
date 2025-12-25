package main

import (
	"encoding/json"
	"net/http"
)

func handlerValidateChirp(w http.ResponseWriter, r *http.Request) {
	type testChirp struct{
		Body string `json:"body"`
	}
	type errorVals struct {
		Error string `json:"error"`
	}
	type returnVals struct {
		Valid bool `json:"valid"`
	}

	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	chirp := testChirp{}
	err := decoder.Decode(&chirp)
	if err != nil {
		newError := errorVals{
			Error: "Something went wrong",
		}
		data, err := json.Marshal(newError)
		if err != nil {
			w.WriteHeader(500)
			w.Write(data)
			return
		}
		w.WriteHeader(500)
		w.Write(data)
		return
	}

	if len(chirp.Body) > 140 {
		newError := errorVals{
			Error: "Chirp is too long",
		}
		data, err := json.Marshal(newError)
		if err != nil {
			w.WriteHeader(500)
			w.Write(data)
			return
		}
		w.WriteHeader(400)
		w.Write(data)
		return
	}

	respBody := returnVals{
		Valid: true,
	}

	data, err := json.Marshal(respBody)
	if err != nil {
		newError := errorVals{
			Error: "Something went wrong",
		}
		data, err := json.Marshal(newError)
		if err != nil {
			w.WriteHeader(500)
			w.Write(data)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}
