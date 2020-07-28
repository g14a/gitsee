package main

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/pretty"
	"gitsee/models"
	"gitsee/service"
)

func main() {
	
	userDetails, _ := service.UserDetails("g14a")
	repoStats := service.GetRepoStats("g14a")
	
	var response models.AbsoluteResponse
	response.UserDetails = userDetails
	response.FrequencyOfLanguages = service.FrequencyOfLanguages(repoStats)
	response.ReposStars = service.RepoStars(repoStats)
	response.ReposForks = service.ReposForks(repoStats)
	bytes, _ := json.Marshal(response)
	
	fmt.Println(string(pretty.Pretty(bytes)))
	
	// r := mux.NewRouter()
	//
	// r.HandleFunc("/user/{username}", service.GetUserInfo)
	//
	// log.Fatal(http.ListenAndServe(":8000", handlers.CORS()(r)))
}
