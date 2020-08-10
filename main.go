package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gitsee/api"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

func main() {
	
	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	
	r.HandleFunc("/user/{username}", api.GetUserInfo)
	r.HandleFunc("/user/{username}/stats/{stat}", api.RepoStats)
	r.HandleFunc("/user/{username}/colorSet", api.GetColorCodes)
	
	if os.Getenv("PORT") == "" {
		log.Fatal(http.ListenAndServe(":8000", handlers.CORS()(handlers.CompressHandler(r))))
	} else {
		if err := http.ListenAndServe(":"+os.Getenv("PORT"), handlers.CORS()(handlers.CompressHandler(r)); err != nil {
			log.Fatal(err)
		}
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		requestDump, err := httputil.DumpRequest(r, true)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(string(requestDump))
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}