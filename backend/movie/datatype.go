package movie

import (
	"encoding/json"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type movieDetails struct {
	Id               string   `json:"id"`
	Poster           string   `json:"poster"`
	Title            string   `json:"title"`
	ReleaseYear      string   `json:"release_year"`
	Tagline          string   `json:"tagline"`
	OriginalLanguage []string `json:"original_language"`
	Popularity       int      `json:"popularity"`
}

type movieErrorResponse struct {
	ErrorMessage string `json:"error"`
}

type listMovieFilters struct {
	genre            *string
	originalLanguage *string
	releaseYear      *int
	searchText       *string
	offset           *int
	limit            *int
}

func normalizeMovies(searchResult redis.FTSearchResult) (result []movieDetails) {
	for _, doc := range searchResult.Docs {
		var originalLanguage []string
		json.Unmarshal([]byte(doc.Fields["original_language"]), &originalLanguage)

		popularityStr := doc.Fields["popularity"]
		popularity, _ := strconv.Atoi(popularityStr)

		result = append(result, movieDetails{
			Id:               doc.Fields["id"],
			Poster:           doc.Fields["poster"],
			Title:            doc.Fields["title"],
			ReleaseYear:      doc.Fields["release_year"],
			Tagline:          doc.Fields["tagline"],
			OriginalLanguage: originalLanguage,
			Popularity:       popularity,
		})
	}
	return
}

type movieClickRequest struct {
	MovieId string `json:"movie_id"`
}
