package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gitsee/color"
	"gitsee/service"
	"net/http"
)

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	username := params["username"]

	var err error

	userDetails, err := service.UserDetails(username)

	if err != nil {
		respondWithError(w, http.StatusNotFound, "Could not get user information")
	}

	respondWithJSON(w, http.StatusOK, userDetails)
}

func RepoStats(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	stat := params["stat"]
	username := params["username"]

	response, ok := service.GetWantedStatFromCache(username, stat)

	// err == nil means the stat is returned from cache if not store it in cache
	if !ok {
		err := service.GetStats(username, 100, 100)
		if err != nil {
			respondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		switch stat {
		case "Contributions":
			respondWithJSON(w, http.StatusOK, service.Contributions)
		case "RepoForks":
			respondWithJSON(w, http.StatusOK, service.ReposForks)
		case "RepoStars":
			respondWithJSON(w, http.StatusOK, service.ReposStars)
		case "RepoLanguages":
			respondWithJSON(w, http.StatusOK, service.LanguageFrequencies)
		case "PrimaryLanguages":
			respondWithJSON(w, http.StatusOK, service.PrimaryLanguages)
		case "PrimaryLanguageStars":
			respondWithJSON(w, http.StatusOK, service.PrimaryLanguageStars)
		}
	} else {
		respondWithJSON(w, http.StatusOK, response)
	}
}

func GetRateLimit(w http.ResponseWriter, r *http.Request)  {
	respondWithJSON(w, http.StatusOK, service.R)
}

func GetColorCodes(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	username := params["username"]

	color.GetColorCodesForLanguages(username, service.PrimaryLanguages)

	respondWithJSON(w, http.StatusOK, color.LanguageColors)
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
