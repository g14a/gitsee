package main

import (
	"github.com/gorilla/mux"
	"gitsee/service"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	
	r.HandleFunc("/user/{username}", service.GetUserInfo)
	
	log.Fatal(http.ListenAndServe(":8000", r))
}
