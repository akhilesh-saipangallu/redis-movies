package movie

// FT.CREATE idx:movies ON JSON PREFIX 1 "movie:" SCHEMA $.title as title TEXT $.genres as genres TAG SEPARATOR "," $.overview as overview TEXT $.release_date as release_date NUMERIC SORTABLE $.tagline as tagline TEXT $.popularity as popularity NUMERIC SORTABLE
const MOVIE_INDEX string = "idx:movies"
