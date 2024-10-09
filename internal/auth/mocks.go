package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type MockAuthenticator struct{}

const testSecret = "test_secret"

var testClaims = jwt.MapClaims{
	"aud": "test-aud",
	"iss": "test-aud",
	"sub": int64(1),
	"exp": time.Now().Add(time.Hour).Unix(),
}

func (m *MockAuthenticator) GenerateToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, testClaims)
	tokenString, _ := token.SignedString([]byte(testSecret))
	return tokenString, nil
}

func (m *MockAuthenticator) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(testSecret), nil
	})
}
