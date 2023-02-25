package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getRes(res http.ResponseWriter, req *http.Request) {
	// fmt.Println(r)
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(movies)
}

func main() {
	var i int = 12
	movies = append(movies, Movie{ID: "1", Title: "Apollo 13", Director: &Director{Firstname: "Christopher", Lastname: "Nolan"}})
	r := mux.NewRouter()
	r.HandleFunc("/", getRes).Methods("GET")
	fmt.Println(i)
	log.Fatal(http.ListenAndServe(":8080", r))
}
