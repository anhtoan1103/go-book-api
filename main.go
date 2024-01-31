package maim

import (
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie

// define the getMovies function to get all movies, we need w to receive the response after processing and need r to send to request. 
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	// if we have the database here, we need to query this from the database.
	// this code will change in the next part.
	json.NewEncoder(w).Encode(movies)
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
	r.HandleFunc("/movie/{id}", deleteMovie).Method("PUT")

	fmt.Println("Starting the server at port 8080!")
	log.Fatal(http.ListenAndServe(":8080", r))
}