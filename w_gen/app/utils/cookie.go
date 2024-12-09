package utils

import (
	"net/http"
	"time"
)

func newHttpOnlyCookie(name, value, path string, t time.Time) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = value
	cookie.Path = path
	cookie.Expires = t
	cookie.HttpOnly = true
	return cookie
}

func NewAccessCookie(value string) *http.Cookie {
	return newHttpOnlyCookie("access", value, "/", time.Now().Add(24*time.Hour))
}

func RemoveAccessCookie() *http.Cookie {
	return newHttpOnlyCookie("access", "", "/", time.Now().Add(-24*time.Hour))
}
