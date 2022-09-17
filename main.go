package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	Id    string `json:"id"`
	Isnb  string `json:"isnb"`
	Title string `json:"title"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var id string = params["id"]
	for _, v := range movies {
		if v.Id == id {
			json.NewEncoder(w).Encode(v)
			return
		}
	}
	http.Error(w, "not found", http.StatusNotFound)
}
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie) // send by reference to update the movie variable
	movies = append(movies, movie)
}
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	for index, v := range movies {
		if v.Id == id {
			movies = append(movies[:index], movies[index+1:]...)
			return
		}
	}
	http.Error(w, "not found", http.StatusNotFound)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var movie Movie
	for k, v := range movies {
		if v.Id == params["id"] {
			movies = append(movies[0:k], movies[k+1:]...)
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.Id = params["id"]
			movies = append(movies, movie)
			return
		}
	}
	http.Error(w, "not found ", http.StatusNotFound)
}
func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{Id: "1", Isnb: "2144", Title: "mad max"})
	movies = append(movies, Movie{Id: "2", Isnb: "1254", Title: "the 100"})
	movies = append(movies, Movie{Id: "3", Isnb: "235", Title: "breaking bad"})
	// add routes
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/createmovie", createMovie).Methods("POST")
	r.HandleFunc("/deletemovie/{id}", deleteMovie).Methods("DELETE")
	r.HandleFunc("/updatemovie/{id}", updateMovie).Methods("PUT")
	// create server
	fmt.Println("server is running on port 8000")
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal("server is down")
	}
}
