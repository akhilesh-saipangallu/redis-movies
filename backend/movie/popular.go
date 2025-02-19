package movie

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlePopularMovies(c *gin.Context) {
	// Get top 10 popular movies
	popularMovies, err := getPopularMovies(c.Request.Context())
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "internal error"})
		return
	}
	c.JSON(http.StatusOK, popularMovies)
}

// func HandlePopularMovies(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")

// 	if r.Method != http.MethodGet {
// 		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	// Get top 10 popular movies
// 	popularMovies, err := getPopularMovies(r.Context())
// 	if err != nil {
// 		log.Println(err)
// 		http.Error(w, newMovieErrorResponseJson("invalid credentials"), http.StatusForbidden)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write([]byte(moviesResponseJson(popularMovies)))
// }
