package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gitsee/api"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	
	r.HandleFunc("/user/{username}", api.GetUserInfo)
	r.HandleFunc("/user/{username}/{stat}", api.RepoStats)
	
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS()(r)))
}
