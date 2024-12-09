package middlewares

import (
	"log"
	"net/http"

	"github.com/m-wilk/w_gen/repository"
	usecase "github.com/m-wilk/w_gen/use-case"
	"github.com/m-wilk/w_gen/utils"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

type AuthMiddleware struct {
	ErrorLog       *log.Logger
	UserRepository repository.UserRepository
	RedisClient    *redis.Client
}

func (a AuthMiddleware) IsLoggedIn(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		access, err := c.Cookie("access")
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, nil)
		}

		token, err := utils.VerifyToken(access.Value)
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				tokenUseCase := usecase.NewToken(a.ErrorLog, a.UserRepository, a.RedisClient)
				result, err := tokenUseCase.Refres(access.Value)
				if err != nil {
					a.ErrorLog.Println(err)
					c.SetCookie(utils.RemoveAccessCookie())
					return echo.NewHTTPError(http.StatusUnauthorized, nil)
				}

				c.SetCookie(utils.NewAccessCookie(result))
				if claims, ok := utils.GetClaim(token); ok {
					c.Set("id", claims["id"])
					c.Set("role", claims["role"])
				}
				return next(c)
			}
			c.SetCookie(utils.RemoveAccessCookie())
			return echo.NewHTTPError(http.StatusUnauthorized, nil)
		}
		if claims, ok := utils.GetClaim(token); ok {
			c.Set("id", claims["id"])
			c.Set("role", claims["role"])
		}

		c.SetCookie(utils.NewAccessCookie(access.Value))
		return next(c)
	}
}
