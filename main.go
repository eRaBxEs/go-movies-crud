package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func main() {

	movies = append(movies, Movie{ID: "1", Isbn: "4563728", Title: "Movie One", Director: &Director{FirstName: "John", LastName: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "4772238", Title: "Movie Two", Director: &Director{FirstName: "Steve", LastName: "Smith"}})
	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Print("Starting server at port: 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {

		if item.ID == params["id"] {
			// delete hack for deleting an element in a slice
			movies = append(movies[:index], movies[index+1:]...)
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	// set the json content type
	w.Header().Set("Content-Type", "application/json")
	// get params
	params := mux.Vars(r)
	// loop over the movies: range
	for index, item := range movies {
		if item.ID == params["id"] {
			// delete the movie
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			// get the movie data from the request
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			// append the new movie to movie slice
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}
