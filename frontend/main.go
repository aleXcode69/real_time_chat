// main.go
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./")))

	addr := ":8080"
	log.Printf("Go server listening on %s …\n", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal(err)
	}
}
