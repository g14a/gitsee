 package service
//
// import (
// 	"encoding/json"
// 	"fmt"
// 	"github.com/gorilla/mux"
// 	"gitsee/models"
// 	"net/http"
// 	"time"
// )
//
// func GetUserInfo(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	username := params["username"]
//
// 	userDetails, err := UserDetails(username)
//
// 	start := time.Now()
// 	repoStats := GetRepoStats(username)
// 	fmt.Println(time.Since(start))
//
// 	var response models.AbsoluteResponse
// 	response.UserDetails = userDetails
//
// //	response.FrequencyOfLanguages = FrequencyOfLanguages(repoStats)
// 	response.ReposStars = RepoStars(repoStats)
// 	response.ReposForks = ReposForks(repoStats)
// 	response.ReposCommits = RepoCommits(repoStats)
//
// 	if err != nil {
// 		respondWithError(w, http.StatusInternalServerError, "Could not get user information")
// 	}
//
// 	respondWithJSON(w, http.StatusOK, response)
// }
//
// func respondWithError(w http.ResponseWriter, httpCode int, message string) {
// 	respondWithJSON(w, httpCode, map[string]string{"error": message})
// }
//
// func respondWithJSON(w http.ResponseWriter, httpCode int, payload interface{}) {
// 	response, _ := json.Marshal(payload)
//
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(httpCode)
//
// 	_, _ = w.Write(response)
// }
