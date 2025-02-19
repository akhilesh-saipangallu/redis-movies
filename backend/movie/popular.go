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
