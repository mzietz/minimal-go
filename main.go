package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type msg struct {
	Msg string `json:"msg"`
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	resp := msg{
		Msg: "Hallo",
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", helloWorld).Methods("GET")
	http.ListenAndServe(":8080", router)
}
