package movie

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/akhilesh-saipangallu/redis-movies/db"
	"github.com/redis/go-redis/v9"
)

func listMoviesWithFilters(ctx context.Context, filters listMovieFilters) ([]movieDetails, error) {
	rdb := db.GetRedisClient()

	var (
		query  string = ""
		offset int    = 0
		limit  int    = 10
	)

	if filters.genre != nil {
		query = query + fmt.Sprintf(`@genres:{"%s"}`, *filters.genre)
	}
	if filters.originalLanguage != nil {
		query = query + fmt.Sprintf(`@original_language:{"%s"}`, *filters.originalLanguage)
	}
	if filters.searchText != nil {
		query = query + fmt.Sprintf(`@title:*%s*`, *filters.searchText)
	}
	// TODO: year filter

	if query == "" {
		query = "*"
	}

	if filters.offset != nil {
		offset = *filters.offset
	}
	if filters.limit != nil {
		limit = *filters.limit
	}

	searchResult, err := rdb.FTSearchWithArgs(
		ctx,
		MOVIE_INDEX,
		query,
		&redis.FTSearchOptions{
			LimitOffset: offset,
			Limit:       limit,
			SortBy:      []redis.FTSearchSortBy{{FieldName: "popularity", Desc: true}},
			Return: []redis.FTSearchReturn{
				{FieldName: "$.id", As: "id"},
				{FieldName: "$.poster", As: "poster"},
				{FieldName: "$.title", As: "title"},
				{FieldName: "$.release_year", As: "release_year"},
				{FieldName: "$.tagline", As: "tagline"},
				{FieldName: "$.original_language", As: "original_language"},
			},
			DialectVersion: 2,
		},
	).Result()

	if err != nil {
		return nil, fmt.Errorf("listMovies: error while running redis query: %w", err)
	}

	if searchResult.Total == 0 {
		return []movieDetails{}, nil
	}

	result := []movieDetails{}
	for _, doc := range searchResult.Docs {
		var originalLanguage []string
		json.Unmarshal([]byte(doc.Fields["original_language"]), &originalLanguage)
		result = append(result, movieDetails{
			Id:               doc.Fields["id"],
			Poster:           doc.Fields["poster"],
			Title:            doc.Fields["title"],
			ReleaseYear:      doc.Fields["release_year"],
			Tagline:          doc.Fields["tagline"],
			OriginalLanguage: originalLanguage,
		})
	}

	return result, nil
}

func getPopularMovies(ctx context.Context) ([]movieDetails, error) {
	rdb := db.GetRedisClient()
	searchResult, err := rdb.FTSearchWithArgs(
		ctx,
		MOVIE_INDEX,
		"*",
		&redis.FTSearchOptions{
			LimitOffset: 0,
			Limit:       10,
			SortBy:      []redis.FTSearchSortBy{{FieldName: "popularity", Desc: true}},
			Return: []redis.FTSearchReturn{
				{FieldName: "$.id", As: "id"},
				{FieldName: "$.poster", As: "poster"},
				{FieldName: "$.title", As: "title"},
				{FieldName: "$.release_year", As: "release_year"},
				{FieldName: "$.tagline", As: "tagline"},
			},
			DialectVersion: 2,
		},
	).Result()

	if err != nil {
		return nil, fmt.Errorf("getPopularMovies: error while running redis query: %w", err)
	}

	if searchResult.Total == 0 {
		return []movieDetails{}, nil
	}

	result := []movieDetails{}
	for _, doc := range searchResult.Docs {
		result = append(result, movieDetails{
			Id:          doc.Fields["id"],
			Poster:      doc.Fields["poster"],
			Title:       doc.Fields["title"],
			ReleaseYear: doc.Fields["release_year"],
			Tagline:     doc.Fields["tagline"],
		})
	}

	return result, nil
}
