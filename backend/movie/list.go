package movie

import (
	"context"
	"log"
	"net/http"
	"strconv"
)

func HandleListMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ctx := context.Background()

	// Get all filters
	filters := extractListMovieFilters(r)

	movies, err := listMoviesWithFilters(ctx, filters)
	if err != nil {
		log.Println(err)
		http.Error(w, newMovieErrorResponseJson("invalid credentials"), http.StatusForbidden)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(moviesResponseJson(movies)))
}

func extractListMovieFilters(r *http.Request) (filters listMovieFilters) {
	var err error
	query := r.URL.Query()

	if genre := query.Get("genre"); genre != "" {
		filters.genre = &genre
	}
	if originalLanguage := query.Get("original_language"); originalLanguage != "" {
		filters.originalLanguage = &originalLanguage
	}

	if searchText := query.Get("search_text"); searchText != "" {
		filters.searchText = &searchText
	}

	releaseYear, err := strconv.Atoi(query.Get("release_year"))
	if err == nil {
		filters.releaseYear = &releaseYear
	}

	offset, err := strconv.Atoi(query.Get("offset"))
	if err == nil {
		filters.offset = &offset
	}

	limit, err := strconv.Atoi(query.Get("limit"))
	if err == nil {
		filters.limit = &limit
	}

	return
}
