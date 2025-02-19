package main

import (
	"log"
	"net/http"

	"github.com/akhilesh-saipangallu/redis-movies/auth"
	"github.com/akhilesh-saipangallu/redis-movies/middlewares"
	"github.com/akhilesh-saipangallu/redis-movies/movie"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.Use(middlewares.CORS)

	router.Handle("POST", "/signup", auth.HandleSignUp)
	router.Handle("POST", "/signin", auth.HandleSignIn)
	router.Handle("GET", "/movies", auth.AuthMiddleware, movie.HandleListMovies)
	router.Handle("GET", "/movies/popular", auth.AuthMiddleware, movie.HandlePopularMovies)
	router.Handle("GET", "/movies/recommendations", auth.AuthMiddleware, movie.HandleRecommendations)
	router.Run()

	log.Println("Server running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Printf("failed to start server, error: %v/n", err)
	}
}
