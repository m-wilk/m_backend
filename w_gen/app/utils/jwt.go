package utils

import (
	"errors"
	"os"
	"strings"
	"time"

	model "github.com/m-wilk/w_gen/models"
	"github.com/golang-jwt/jwt"
)

func GenerateJWT(user model.User, exp time.Duration) (string, error) {
	secretKey := os.Getenv("SECRET_KEY")
	cleanSecretKey := strings.TrimSpace(secretKey)
	if cleanSecretKey == "" && len(cleanSecretKey) < 64 {
		return "", errors.New("invalid secret key")
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"iss":   "user service",
		"role":  user.Role,
		"exp":   time.Now().Add(exp).Unix(),
		"iat":   time.Now().Unix(),
	})

	return claims.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
}

func ShouldTokenRefresh(tokenError error) bool {
	if ve, ok := tokenError.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorExpired != 0 {
			return true
		}
	}

	return false
}

func GetClaim(token *jwt.Token) (jwt.MapClaims, bool) {
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, ok
	}
	return jwt.MapClaims{}, false
}
