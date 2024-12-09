package handlers

import (
	"fmt"
	"net/http"

	"github.com/m-wilk/w_gen/repository"
	usecase "github.com/m-wilk/w_gen/use-case"
	"github.com/m-wilk/w_gen/utils"
	"github.com/labstack/echo/v4"
)

func (h *Handler) UserDetail(c echo.Context) error {
	userId := c.Get("id")
	if userId == "" {
		return c.JSON(http.StatusBadRequest, "no user found")
	}
	result, err := h.Core.Repository.UserRepository.FindOne(repository.UserQuery{ID: fmt.Sprint(userId)})
	if err != nil {
		h.Core.ErrorLog.Println(err)
		echo.NewHTTPError(http.StatusBadRequest, "get users error")
	}
	return c.JSON(http.StatusOK, result)
}

func (h *Handler) UsersList(c echo.Context) error {
	userId := c.Get("id")
	if userId == "" {
		return c.JSON(http.StatusBadRequest, "no user found")
	}
	result, err := h.Core.Repository.UserRepository.FindOne(repository.UserQuery{ID: fmt.Sprint(userId)})
	if err != nil {
		h.Core.ErrorLog.Println(err)
		echo.NewHTTPError(http.StatusBadRequest, "get users error")
	}
	return c.JSON(http.StatusOK, result)
}

type BaseAuthParams struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (h *Handler) Login(c echo.Context) error {
	ap := new(BaseAuthParams)
	if err := c.Bind(ap); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	dv := NewValidator()
	if err := dv.Validate(ap); err != nil {
		return err
	}

	loginUseCase := usecase.NewLogin(h.Core.ErrorLog, h.Core.Repository.UserRepository, h.Core.RedisClient)
	result, err := loginUseCase.Base(ap.Email, ap.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	c.SetCookie(utils.NewAccessCookie(result))
	return c.JSON(http.StatusOK, loginUseCase.User)
}

func (h *Handler) Logout(c echo.Context) error {
	userId := c.Get("id")
	if userId == "" {
		return c.JSON(http.StatusBadRequest, "logged out error")
	}
	logoutUseCase := usecase.NewLogout(h.Core.RedisClient)
	logoutUseCase.Base(fmt.Sprint(userId))
	c.SetCookie(utils.NewAccessCookie(""))
	return c.JSON(http.StatusOK, "logged out")
}

func (h *Handler) Register(c echo.Context) error {
	ap := new(BaseAuthParams)
	if err := c.Bind(ap); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	dv := NewValidator()
	if err := dv.Validate(ap); err != nil {
		return err
	}

	registerUseCase := usecase.NewRegister(h.Core.ErrorLog, h.Core.Repository.UserRepository, h.Core.RedisClient)
	result, err := registerUseCase.Base(ap.Email, ap.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, result)
}
