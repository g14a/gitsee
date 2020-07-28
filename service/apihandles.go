package service

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	username := params["username"]

	userDetails, err := UserDetails(username)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not get user information")
	}

	respondWithJSON(w, http.StatusOK, userDetails)
}



func respondWithError(w http.ResponseWriter, httpCode int, message string) {
	respondWithJSON(w, httpCode, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, httpCode int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)

	_, _ = w.Write(response)
}
