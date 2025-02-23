package movie

import (
	"log"
	"net/http"

	"github.com/akhilesh-saipangallu/redis-movies/auth"
	"github.com/gin-gonic/gin"
)

func HandleRecommendations(c *gin.Context) {
	// Get user id
	currentUserId, ok := auth.GetCurrentUserId(c)
	if !ok {
		log.Println("error: missing user id in request context")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
	}

	// Get stored list of movies for recommendations
	recommendedMovieIds := getStoredListOfRecommendedMovieIds(c.Request.Context(), currentUserId)

	// List all the movie details of recommended movies
	movies, err := getMovieDetails(c.Request.Context(), recommendedMovieIds)
	if err != nil {
		log.Println("error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}
	c.JSON(http.StatusOK, movies)
}
