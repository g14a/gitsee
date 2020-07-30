package main

import (
	"gitsee/service"
)

func main() {

	//	service.UserDetails("g14a")

	service.ForksStarsLanguages("kovidgoyal", 100, 100)

	// r := mux.NewRouter()
	//
	// r.HandleFunc("/user/{username}", service.GetUserInfo)
	//
	// log.Fatal(http.ListenAndServe(":8000", handlers.CORS()(r)))
}
