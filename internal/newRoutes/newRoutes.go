package newroutes

import (
	"fmt"
	"net/http"
)

func NewController(Mux *http.Handler) {
	Mux.HandleFunc("api/weather", weatherController)
}

func weatherController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Weather Controller!")
}
