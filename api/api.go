package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gitsee/service"
	"net/http"
	"sync"
)

var once sync.Once

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	username := params["username"]

	var err error
	
	userDetails, err := service.UserDetails(username)
	
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not get user information")
	}

	respondWithJSON(w, http.StatusOK, userDetails)
}

func RepoStats(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)
	stat := params["stat"]
	username := params["username"]
	
	service.GetAllStats(username, 100, 100)
	
	switch stat {
	case "repoForks":
		respondWithJSON(w, http.StatusOK, service.ReposForks)
	case "repoStars":
		respondWithJSON(w, http.StatusOK, service.ReposStars)
	case "repoLanguages":
		respondWithJSON(w, http.StatusOK, service.LanguageFrequencies)
	case "primaryLanguages":
		respondWithJSON(w, http.StatusOK, service.PrimaryLanguages)
	case "primaryLanguageStars":
		respondWithJSON(w, http.StatusOK, service.PrimaryLanguageStars)
	}
}

func ReposForks(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, service.ReposForks)
}

func ReposStars(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, service.ReposStars)
}

func LanguageFrequencies(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, service.LanguageFrequencies)
}

func PrimaryLanguages(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, service.PrimaryLanguages)
}

func PrimaryLanguagesStars(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, service.PrimaryLanguageStars)
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
