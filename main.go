package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gitsee/api"
	"log"
	"net/http"
	"os"
)

func main() {

	r := mux.NewRouter()

	r.Handle("/user/{username}",
		handlers.LoggingHandler(os.Stdout, http.HandlerFunc(api.GetUserInfo)))
	r.Handle("/user/{username}/stats/{stat}",
		handlers.LoggingHandler(os.Stdout, http.HandlerFunc(api.RepoStats)))
	r.Handle("/user/{username}/colorSet",
		handlers.LoggingHandler(os.Stdout, http.HandlerFunc(api.GetColorCodes)))

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS()(handlers.CompressHandler(r))))
}
