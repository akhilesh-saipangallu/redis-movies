package movie

import "encoding/json"

func newMovieErrorResponseJson(errorMessage string) string {
	authError := movieErrorResponse{ErrorMessage: errorMessage}
	authErrorJson, _ := json.Marshal(authError)
	return string(authErrorJson)
}

func moviesResponseJson(movies []movieDetails) string {
	responseJson, _ := json.Marshal(movies)
	return string(responseJson)
}
