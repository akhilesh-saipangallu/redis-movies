package movie

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HandleListMovies(c *gin.Context) {
	// Get all filters
	filters := extractListMovieFilters(c)

	movies, err := listMoviesWithFilters(c.Request.Context(), filters)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	// if filters.searchText != nil {
	// 	go func() {
	// 		userId, ok := auth.GetCurrentUserId(c)
	// 		if !ok {
	// 			log.Println("error: HandleListMovies: failed to get user id from context to track movies")
	// 		}
	// 		err = trackUserSearch(context.Background(), userId, movies)
	// 		if err != nil {
	// 			log.Println("error: HandleListMovies:", err)
	// 		}
	// 	}()
	// }

	c.JSON(http.StatusOK, movies)
}

func extractListMovieFilters(c *gin.Context) (filters listMovieFilters) {
	var err error

	if genre := c.DefaultQuery("genre", ""); genre != "" {
		filters.genre = &genre
	}
	if originalLanguage := c.DefaultQuery("original_language", ""); originalLanguage != "" {
		filters.originalLanguage = &originalLanguage
	}

	if searchText := c.DefaultQuery("search_text", ""); searchText != "" {
		filters.searchText = &searchText
	}

	releaseYear, err := strconv.Atoi(c.DefaultQuery("release_year", ""))
	if err == nil {
		filters.releaseYear = &releaseYear
	}

	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err == nil {
		filters.offset = &offset
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err == nil {
		filters.limit = &limit
	}

	return
}
