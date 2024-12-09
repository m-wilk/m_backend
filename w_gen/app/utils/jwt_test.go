package utils

import (
	"os"
	"testing"
	"time"

	model "github.com/m-wilk/w_gen/models"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func Test_GenerateJWT(t *testing.T) {
	testSecretKey := "test_secret_key"
	os.Setenv("SECRET_KEY", testSecretKey)

	user := model.User{
		ID:    "1",
		Email: "test-email@example.com",
		Role:  model.SuperAdminRole,
	}

	expiredDuration := time.Minute * 15

	tokenString, err := GenerateJWT(user, expiredDuration)
	assert.NoError(t, err, "Expected no error while generating token")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(testSecretKey), nil
	})
	assert.NoError(t, err, "Expected no error while parsing token")
	assert.True(t, token.Valid, "Expected token to be valid")

	claims, ok := token.Claims.(jwt.MapClaims)
	assert.True(t, ok, "Expected claims to be of type jwt.MapClaims")
	assert.Equal(t, user.ID, claims["id"], "Expected user ID to be in claims")
	assert.Equal(t, user.Email, claims["email"], "Expected user email to be in claims")
	assert.Equal(t, "user service", claims["iss"], "Expected issuer to be 'user service'")
	assert.Equal(t, string(model.SuperAdminRole), claims["role"], "Expected audience to be user role (super-admin)")
	assert.WithinDuration(t, time.Now().Add(expiredDuration), time.Unix(int64(claims["exp"].(float64)), 0), time.Second, "Expected expiration time to be within the set duration")
	assert.WithinDuration(t, time.Now(), time.Unix(int64(claims["iat"].(float64)), 0), time.Second, "Expected issued at time to be current time")
}

func Test_GenerateJWTInvalidSecret(t *testing.T) {
	os.Setenv("SECRET_KEY", "")

	user := model.User{
		ID:    "1",
		Email: "test-email@example.com",
		Role:  model.SuperAdminRole,
	}

	expiredDuration := time.Minute * 15

	_, err := GenerateJWT(user, expiredDuration)
	assert.Error(t, err, "Expected error while generating token with empty secret key")

	os.Setenv("SECRET_KEY", "\t\n  ")
	_, err = GenerateJWT(user, expiredDuration)
	assert.Error(t, err, "Expected error while generating token with empty secret key")
}

func Test_VerifyToken(t *testing.T) {
	testSecretKey := "test_secret_key"
	os.Setenv("SECRET_KEY", testSecretKey)

	user := model.User{
		ID:    "1",
		Email: "test-email@example.com",
		Role:  model.SuperAdminRole,
	}

	validToken, _ := GenerateJWT(user, time.Minute*15)
	invalidToken := "invalidToken"

	os.Setenv("SECRET_KEY", "differentSecret")
	invalidSignatureToken, _ := GenerateJWT(user, time.Minute*15)
	os.Setenv("SECRET_KEY", testSecretKey)

	tests := []struct {
		name   string
		token  string
		isErr  bool
		errMsg string
	}{
		{"Valid token", validToken, false, ""},
		{"Invalid token", invalidToken, true, "token contains an invalid number of segments"},
		{"Wrong signature", invalidSignatureToken, true, "signature is invalid"},
	}

	for _, tt := range tests {
		token, err := VerifyToken(tt.token)

		if tt.isErr {
			if err == nil {
				t.Error("Expected error, get nil")
			} else if err.Error() != tt.errMsg {
				t.Errorf("VerifyToken() error = %v, expected error to contain: %v", err, tt.errMsg)
			}
		}

		if err == nil && !token.Valid {
			t.Errorf("Expected token to be valid, but it was invalid")
		}
	}
}
