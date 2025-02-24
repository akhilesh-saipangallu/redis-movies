package movie

import (
	"context"
	"log"
	"net/http"

	"github.com/akhilesh-saipangallu/redis-movies/auth"
	"github.com/gin-gonic/gin"
)

func HandleMovieClickEvent(c *gin.Context) {
	// Get user id
	currentUserId, ok := auth.GetCurrentUserId(c)
	if !ok {
		log.Println("error: missing user id in request context")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
	}

	var requestData movieClickRequest
	err := c.BindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	go func(userId string, movieId string) {
		err := trackUserClick(context.Background(), userId, movieId)
		if err != nil {
			log.Println("error: HandleListMovies:", err)
		}
	}(currentUserId, requestData.MovieId)

	c.JSON(http.StatusOK, gin.H{"success": true})
}
