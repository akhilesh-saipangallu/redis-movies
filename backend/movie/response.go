package movie

import "encoding/json"

func newMovieErrorResponseJson(errorMessage string) string {
	errorResponse := movieErrorResponse{ErrorMessage: errorMessage}
	errorJson, _ := json.Marshal(errorResponse)
	return string(errorJson)
}

func moviesResponseJson(movies []movieDetails) string {
	responseJson, _ := json.Marshal(movies)
	return string(responseJson)
}
