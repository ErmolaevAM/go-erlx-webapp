package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := createNewRouter()
	http.ListenAndServe(":8080", router)
}

func createNewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/hello", handlerWelcome).Methods("GET")
	return router
}

func handlerWelcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "It's alive!")
}
