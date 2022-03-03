package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/ping", pingpongHandler).Methods("GET")
	http.ListenAndServe(":8080", r)
}

func pingpongHandler(w http.ResponseWriter, r *http.Request) {
	resp := struct {
		Message string `json:"message"`
	}{
		Message: "pong",
	}

	json.NewEncoder(w).Encode(&resp)

}
