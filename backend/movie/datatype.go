package movie

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
