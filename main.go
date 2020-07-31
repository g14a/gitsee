package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gitsee/api"
	"log"
	"net/http"
)

func main() {

	// cache.RistrettoCache.Set("name", "gowtham", 1)
	// time.Sleep(10*time.Millisecond)
	// value, found := cache.RistrettoCache.Get("name")
	// if !found {
	// 	fmt.Println("missing")
	// }
	// fmt.Println(value)

	r := mux.NewRouter()

	r.HandleFunc("/user/{username}", api.GetUserInfo)
	r.HandleFunc("/user/{username}/{stat}", api.RepoStats)

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS()(r)))
}
