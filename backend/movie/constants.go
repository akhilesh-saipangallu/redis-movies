package movie

// FT.CREATE idx:movies ON JSON PREFIX 1 "movie:" SCHEMA $.id as id TAG $.title as title TEXT $.genres as genres TAG SEPARATOR "," $.overview as overview TEXT $.release_year as release_year NUMERIC SORTABLE $.tagline as tagline TEXT $.popularity as popularity NUMERIC SORTABLE $.original_language as original_language TAG SEPARATOR ","  $.embedding as embedding VECTOR FLAT 6 TYPE FLOAT32 DIM 768 DISTANCE_METRIC COSINE
// FT.CREATE idx:movies ON JSON PREFIX 1 "movie:" SCHEMA
//     $.id as id TAG
//     $.title as title TEXT
//     $.genres as genres TAG SEPARATOR ","
//     $.overview as overview TEXT
//     $.release_year as release_year NUMERIC SORTABLE
//     $.tagline as tagline TEXT
//     $.popularity as popularity NUMERIC SORTABLE
//     $.original_language as original_language TAG SEPARATOR "," 
//     $.embedding as embedding VECTOR FLAT 6 TYPE FLOAT32 DIM 768 DISTANCE_METRIC COSINE

const MOVIE_INDEX string = "idx:movies"
const USER_RECOMMENDATIONS_DOC_PREFIX string = "user:recommendations:"
