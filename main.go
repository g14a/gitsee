package main

import (
	"github.com/gorilla/mux"
	"gitsee/service"
	"log"
	"net/http"
	"github.com/gorilla/handlers"
)

func main() {
	r := mux.NewRouter()
	
	r.HandleFunc("/user/{username}", service.GetUserInfo)
	
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS()(r)))
}
