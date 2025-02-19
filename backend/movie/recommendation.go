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

// func HandleRecommendations(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Access-Control-Allow-Origin", "*")

// 	if r.Method != http.MethodGet {
// 		http.Error(w, newMovieErrorResponseJson("method not allowed"), http.StatusMethodNotAllowed)
// 		return
// 	}

// 	// Get user id
// 	currentUserId, ok := auth.GetCurrentUserId(r)
// 	if !ok {
// 		http.Error(w, newMovieErrorResponseJson("bad request"), http.StatusBadRequest)
// 	}

// 	// Get stored list of movies for recommendations
// 	recommendedMovieIds := getStoredListOfRecommendedMovieIds(r.Context(), currentUserId)

// 	// List all the movie details of recommended movies
// 	movies, err := getMovieDetails(r.Context(), recommendedMovieIds)
// 	if err != nil {
// 		log.Println(err)
// 		http.Error(w, newMovieErrorResponseJson("internal error"), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write([]byte(moviesResponseJson(movies)))
// }
