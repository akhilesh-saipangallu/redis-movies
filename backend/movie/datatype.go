package movie

type movieDetails struct {
	Id          string `json:"id"`
	Poster      string `json:"poster"`
	Title       string `json:"title"`
	ReleaseDate string `json:"release_date"`
	VoteAverage string `json:"vote_average"`
}

type movieErrorResponse struct {
	ErrorMessage string `json:"error"`
}

type listMovieFilters struct {
	genre       *string
	releaseYear *int
	searchText  *string
	offset      *int
	limit       *int
}
