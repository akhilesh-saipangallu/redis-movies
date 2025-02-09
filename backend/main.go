package main

import (
	"log"
	"net/http"

	"github.com/akhilesh-saipangallu/redis-movies/auth"
	"github.com/akhilesh-saipangallu/redis-movies/movie"
)

func main() {
	http.HandleFunc("/signup", auth.HandleSignUp)
	http.HandleFunc("/signin", auth.HandleSignIn)

	// movies
	http.HandleFunc("/movies", movie.HandleListMovies)
	http.HandleFunc("/movies/popular", movie.HandlePopularMovies)

	log.Println("Server running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Printf("failed to start server, error: %v/n", err)
	}
}
