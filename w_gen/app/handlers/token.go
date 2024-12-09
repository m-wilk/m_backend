package handlers

import (
	"net/http"
	"strings"

	usecase "github.com/m-wilk/w_gen/use-case"
	"github.com/m-wilk/w_gen/utils"
	"github.com/labstack/echo/v4"
)

func (h *Handler) RefreshToken(c echo.Context) error {
	accessCookie, err := c.Cookie("access")
	if err != nil || strings.TrimSpace(accessCookie.Value) == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, nil)
	}

	tokenUseCase := usecase.NewToken(h.Core.ErrorLog, h.Core.Repository.UserRepository, h.Core.RedisClient)
	result, err := tokenUseCase.Refres(accessCookie.Value)

	if err != nil {
		h.Core.ErrorLog.Println(err)
		return echo.NewHTTPError(http.StatusUnauthorized, nil)
	}

	c.SetCookie(utils.NewAccessCookie(result))
	return c.JSON(http.StatusOK, nil)
}

func (h *Handler) VerifyToken(c echo.Context) error {
	accessCookie, err := c.Cookie("access")
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "no token to verify")
	}
	_, err = utils.VerifyToken(accessCookie.Value)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "verify token error")
	}

	return c.JSON(http.StatusOK, nil)
}
