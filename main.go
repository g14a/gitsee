package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gitsee/service"
	"log"
	"net/http"
)

func main() {
	
	// userDetails, _ := service.UserDetails("kovidgoyal")
	// repoStats := service.GetRepoStats("kovidgoyal")
	//
	// var response models.AbsoluteResponse
	// response.UserDetails = userDetails
	// response.FrequencyOfLanguages = service.FrequencyOfLanguages(repoStats)
	// response.ReposStars = service.RepoStars(repoStats)
	// response.ReposForks = service.ReposForks(repoStats)
	// bytes, _ := json.Marshal(response)
	
//	fmt.Println(string(pretty.Pretty(bytes)))
	
	r := mux.NewRouter()
	
	r.HandleFunc("/user/{username}", service.GetUserInfo)
	
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS()(r)))
}
