package movie

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

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
		query = query + fmt.Sprintf(`@title:%%%s%%`, *filters.searchText)
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
			// SortBy:      []redis.FTSearchSortBy{{FieldName: "popularity", Desc: true}},
			SortBy: []redis.FTSearchSortBy{{FieldName: "release_year", Desc: true}},
			Return: []redis.FTSearchReturn{
				{FieldName: "$.id", As: "id"},
				{FieldName: "$.poster", As: "poster"},
				{FieldName: "$.title", As: "title"},
				{FieldName: "$.release_year", As: "release_year"},
				{FieldName: "$.tagline", As: "tagline"},
				{FieldName: "$.original_language", As: "original_language"},
				{FieldName: "$.popularity", As: "popularity"},
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

	result := normalizeMovies(searchResult)
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
				{FieldName: "$.original_language", As: "original_language"},
				{FieldName: "$.popularity", As: "popularity"},
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

	result := normalizeMovies(searchResult)
	return result, nil
}

func getStoredListOfRecommendedMovieIds(ctx context.Context, userId string) []string {
	rdb := db.GetRedisClient()
	commandResult, err := rdb.JSONGet(ctx, fmt.Sprintf("%s%s", USER_RECOMMENDATIONS_DOC_PREFIX, userId), "$.[*].id").Result()
	if err != nil {
		return []string{}
	}

	movieIds := []string{}
	err = json.Unmarshal([]byte(commandResult), &movieIds)
	if err != nil {
		log.Println("error unmarshaling string:", err)
	}

	return movieIds
}

func getMovieDetails(ctx context.Context, movieIds []string) (result []movieDetails, err error) {
	if len(movieIds) == 0 {
		return
	}

	rdb := db.GetRedisClient()
	searchResult, err := rdb.FTSearchWithArgs(
		ctx,
		MOVIE_INDEX,
		fmt.Sprintf("@id:{%s}", strings.Join(movieIds, "|")),
		&redis.FTSearchOptions{
			Return: []redis.FTSearchReturn{
				{FieldName: "$.id", As: "id"},
				{FieldName: "$.poster", As: "poster"},
				{FieldName: "$.title", As: "title"},
				{FieldName: "$.release_year", As: "release_year"},
				{FieldName: "$.tagline", As: "tagline"},
				{FieldName: "$.original_language", As: "original_language"},
				{FieldName: "$.popularity", As: "popularity"},
			},
			DialectVersion: 2,
		},
	).Result()

	if err != nil {
		err = fmt.Errorf("getPopularMovies: error while running redis query: %w", err)
		return
	}

	if searchResult.Total == 0 {
		return []movieDetails{}, nil
	}

	result = normalizeMovies(searchResult)
	return result, nil
}

// func trackUserSearch(ctx context.Context, userId string, movies []movieDetails) error {
// 	movieIds := []string{}
// 	for i, movie := range movies {
// 		if i >= 2 {
// 			break
// 		}
// 		movieIds = append(movieIds, movie.Id)
// 	}

// 	if len(movieIds) == 0 {
// 		log.Println("warning: no movies to track for user:", userId)
// 		return nil
// 	}

// 	movieIdsStr, _ := json.Marshal(movieIds)

// 	rdb := db.GetRedisClient()
// 	_, err := rdb.XAdd(
// 		ctx, &redis.XAddArgs{
// 			Stream: STREAM_USER_SEARCH_HISTORY,
// 			Values: map[string]interface{}{
// 				"user_id":   userId,
// 				"movie_ids": movieIdsStr,
// 			},
// 		},
// 	).Result()

// 	if err != nil {
// 		return fmt.Errorf("trackUserSearch: %w", err)
// 	}
// 	log.Println("trackUserSearch: successful for user:", userId)
// 	return nil
// }

func trackUserClick(ctx context.Context, userId string, movieId string) error {
	movieIds := []string{movieId}
	movieIdsStr, _ := json.Marshal(movieIds)

	rdb := db.GetRedisClient()
	_, err := rdb.XAdd(
		ctx, &redis.XAddArgs{
			Stream: STREAM_USER_SEARCH_HISTORY,
			Values: map[string]interface{}{
				"user_id":   userId,
				"movie_ids": movieIdsStr,
			},
		},
	).Result()

	if err != nil {
		return fmt.Errorf("trackUserClick: %w", err)
	}
	log.Println("trackUserClick: successful for user:", userId)
	return nil
}
