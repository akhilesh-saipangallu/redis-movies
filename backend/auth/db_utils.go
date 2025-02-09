package auth

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/akhilesh-saipangallu/redis-movies/db"
	"github.com/akhilesh-saipangallu/redis-movies/utils"
	"github.com/redis/go-redis/v9"
)

func doesUserExists(ctx context.Context, email string) (bool, error) {
	rdb := db.GetRedisClient()
	query := fmt.Sprintf(`@email:{"%s"}`, email)

	aggResult, err := rdb.FTAggregateWithArgs(
		ctx,
		USER_INDEX,
		query,
		&redis.FTAggregateOptions{
			GroupBy: []redis.FTAggregateGroupBy{
				{
					Reduce: []redis.FTAggregateReducer{
						{
							Reducer: redis.SearchCount,
							As:      "userCount",
						},
					},
				},
			},
			DialectVersion: 2,
		},
	).Result()

	if err != nil {
		return false, fmt.Errorf("doesUserExists: error while running redis query: %w", err)
	}

	if aggResult == nil || aggResult.Total != 1 {
		log.Println("bad result: ", aggResult)
		return false, nil
	}

	userCount := aggResult.Rows[0].Fields["userCount"]
	userCountInt, err := strconv.ParseInt(userCount.(string), 10, 64)
	if err != nil {
		return false, fmt.Errorf("doesUserExists: error converting userCount: %w", err)
	}

	if userCountInt > 0 {
		return true, nil
	}

	return false, nil
}

func createUser(ctx context.Context, signUpRequest SignUpRequest) error {
	hashedPassword, err := hashPassword(signUpRequest.Password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}

	userId := utils.GenerateUUID()
	user := User{
		Id:        userId,
		Email:     signUpRequest.Email,
		FirstName: signUpRequest.FirstName,
		LastName:  signUpRequest.LastName,
		Password:  hashedPassword,
	}

	rdb := db.GetRedisClient()
	_, err = rdb.JSONSet(ctx, fmt.Sprintf("user:%s", userId), "$", user).Result()
	if err != nil {
		return fmt.Errorf("failed to create user in redis: %v", err)
	}

	return nil
}

func getUserDetails(ctx context.Context, email string) (*User, error) {
	rdb := db.GetRedisClient()
	query := fmt.Sprintf(`@email:{"%s"}`, email)

	searchResult, err := rdb.FTSearchWithArgs(
		ctx,
		USER_INDEX,
		query,
		&redis.FTSearchOptions{
			Return:         []redis.FTSearchReturn{
				{FieldName: "$.id", As: "id"},
				{FieldName: "$.email", As: "email"},
				{FieldName: "$.password", As: "password"},
			},
			DialectVersion: 2,
		},
	).Result()

	if err != nil {
		return nil, fmt.Errorf("getUserDetails: error while running redis query: %w", err)
	}

	if searchResult.Total == 0 {
		return nil, fmt.Errorf("getUserDetails: no user found with email %s", email)
	}

	if searchResult.Total > 1 {
		return nil, fmt.Errorf("getUserDetails: more than one user with email %s", email)
	}

	user := User{
		Id: searchResult.Docs[0].Fields["id"],
		Email: searchResult.Docs[0].Fields["email"],
		Password: searchResult.Docs[0].Fields["password"],
	}
	return &user, nil
}
