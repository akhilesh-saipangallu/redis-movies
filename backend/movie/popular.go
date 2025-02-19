package movie

import (
	"log"
	"net/http"
)

func HandlePopularMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get top 10 popular movies
	popularMovies, err := getPopularMovies(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, newMovieErrorResponseJson("invalid credentials"), http.StatusForbidden)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(moviesResponseJson(popularMovies)))
}
