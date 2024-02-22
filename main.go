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
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

// define the getMovies function to get all movies, we need w to receive the response after processing and need r to send to request.
func getMovies(w http.ResponseWriter, r *http.Request) {
	// set the header Content-type to point out the type of the data is application or json.
	w.Header().Set("Content-type", "application/json")
	// if we have the database here, we need to query this from the database.
	// this code will change in the next part.
	// this will encode the movies and then use NewEncoder to write to w and use json to return to client.
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	// set the header Content-type to point out the type of the data is application or json.
	w.Header().Set("Content-type", "application/json")
	// get the list params from the url.
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	// set the header Content-type to point out the type of the data is application or json.
	w.Header().Set("Content-type", "application/json")
	// get the list params from the url.
	params := mux.Vars(r)

	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			break
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	// set the header Content-type to point out the type of the data is application or json.
	w.Header().Set("Content-type", "application/json")
	// declare a movie variable as Movie type
	var movie Movie
	// receive data from body and decode it to movie variable.
	_ = json.NewDecoder(r.Body).Decode(&movie)
	// generate a random ID for the new movie that receive from body.
	// if the movieId exists?
	movie.ID = strconv.Itoa(rand.Intn(100000000))

	// append new movie to movies
	movies = append(movies, movie)
	// return the new movie to the response from view the result.
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	// set the header Content-type to point out the type of the data is application or json.
	w.Header().Set("Content-type", "application/json")
	// get all params from the request.
	params := mux.Vars(r)
	// loop to all the element of movies to find the ID.
	for index, item := range movies {
		for item.ID == params["id"] {
			// remove the movie with the same id
			movies = append(movies[:index], movies[index+1:]...)
			// add new movie
			// create a new movie variable
			var movie Movie
			// receive data from request body and decode it to movie variable
			_ = json.NewDecoder(r.Body).Decode(&movie)
			// generate a new ID for the new movie
			movie.ID = params["id"]
			// append new movie to movies
			movies = append(movies, movie)
			// send data to responseWriter to response the updated movies
			json.NewEncoder(w).Encode(movies)
			return
		}
	}
}

func main() {
	// use mux library to define the router for the handler.
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "123483", Title: "cuon theo chieu gio", Director: &Director{Firstname: "toto", Lastname: "beo"}})
	movies = append(movies, Movie{ID: "2", Isbn: "343423", Title: "cuon theo",
		Director: &Director{Firstname: "sora", Lastname: "map"}})
	// route to /movies to getAll movies.
	r.HandleFunc("/movies", getMovies).Methods("GET")
	// route to /movie/{id} to getMovieById.
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	// route to /movie to createMovie.
	r.HandleFunc("/movie", createMovie).Methods("POST")
	// route to /movie/{id} to updateMovie.
	r.HandleFunc("/movie/{id}", updateMovie).Methods("PUT")
	// route to /movie/{id} to deleteMovie.
	r.HandleFunc("/movie/{id}", deleteMovie).Methods("PUT")

	fmt.Println("Starting the server at port 8080!")
	log.Fatal(http.ListenAndServe(":8080", r))
}
