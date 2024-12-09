package usecase

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/m-wilk/w_gen/constants"
	"github.com/m-wilk/w_gen/repository"
	"github.com/m-wilk/w_gen/utils"
	"github.com/redis/go-redis/v9"
)

type Token interface {
	Refres(accessToken string) (string, error)
}

type token struct {
	ErrorLog       *log.Logger
	UserRepository repository.UserRepository
	RedisClient    *redis.Client
}

func NewToken(logger *log.Logger, userRepository repository.UserRepository, redis *redis.Client) token {
	return token{ErrorLog: logger, UserRepository: userRepository, RedisClient: redis}
}

func (t *token) Refres(accessToken string) (string, error) {
	baseError := errors.New("refresh token error")
	// 1. get token struct and it's claim (need user id)
	token, _ := utils.VerifyToken(accessToken)
	claim, ok := utils.GetClaim(token)
	if !ok {
		return "", errors.New("get claim error")
	}

	// 2. pick and verify user refresh token
	userID := fmt.Sprint(claim["id"])
	refreshToken, err := t.RedisClient.Get(context.Background(), userID).Result()
	if err != nil {
		return "", errors.New("logged out")
	}

	_, err = utils.VerifyToken(refreshToken)
	if err != nil {
		return "", baseError
	}

	// 3. If we don't have any error from verify refresh token then we generate new access token
	// TODO don't get data from DB just get this from claim
	user, err := t.UserRepository.FindOne(repository.UserQuery{ID: userID})
	if err != nil {
		t.ErrorLog.Println(err)
		return "", baseError
	}

	newAccessToken, err := utils.GenerateJWT(*user, constants.ACCESS_TOKEN_EXP)
	if err != nil {
		t.ErrorLog.Println(err)
		return "", baseError
	}

	return newAccessToken, nil
}
