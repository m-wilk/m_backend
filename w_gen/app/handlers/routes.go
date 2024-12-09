package handlers

import (
	"net/http"

	guard "github.com/m-wilk/w_gen/handlers/guards"
	"github.com/m-wilk/w_gen/handlers/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (h *Handler) Routes(e *echo.Echo) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	apiGroup := e.Group("/api/v1")

	apiGroup.GET("/creative-staff", h.CreativeStaff)

	apiGroup.POST("/login", h.Login)
	apiGroup.POST("/register", h.Register)

	apiGroup.GET("/token/verify", h.VerifyToken)
	apiGroup.GET("/token/refresh", h.RefreshToken)

	authMiddleware := middlewares.AuthMiddleware{
		ErrorLog:       h.Core.ErrorLog,
		UserRepository: h.Core.Repository.UserRepository,
		RedisClient:    h.Core.RedisClient,
	}

	apiGroup.Use(authMiddleware.IsLoggedIn)

	apiGroup.GET("/logout", h.Logout)
	apiGroup.GET("/user", h.UserDetail)
	apiGroup.GET("/users", guard.AdminAuthGuard(h.UsersList))
}
