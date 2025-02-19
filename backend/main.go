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
	http.Handle("/movies", auth.AuthMiddleware(http.HandlerFunc(movie.HandleListMovies)))
	http.Handle("/movies/popular", auth.AuthMiddleware(http.HandlerFunc(movie.HandlePopularMovies)))
	http.Handle("/movies/recommendations", auth.AuthMiddleware(http.HandlerFunc(movie.HandleRecommendations)))

	log.Println("Server running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Printf("failed to start server, error: %v/n", err)
	}
}
