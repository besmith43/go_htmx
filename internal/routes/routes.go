package routes

import (
	"fmt"
	"net/http"
)

// for public funcs and vars use a capital letter to start the name
// for private funcs and vars use a lower case letter to start the name

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/api/data", apiDataHandler)

	return mux
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the home page!")
}

func apiDataHandler(w http.ResponseWriter, r *http.Request) {
	data := "Some Data from the API"
	fmt.Fprintln(w, data)
}
