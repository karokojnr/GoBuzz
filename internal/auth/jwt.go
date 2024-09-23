package auth

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type JWTAuthenticator struct {
	secret   string
	audience string
	issuer   string
}

func NewJWTAuthenticator(secret, audience, issuer string) *JWTAuthenticator {
	return &JWTAuthenticator{secret: secret, audience: audience, issuer: issuer}
}

func (a *JWTAuthenticator) GenerateToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(a.secret))
}

func (a *JWTAuthenticator) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(a.secret), nil
	},
		jwt.WithExpirationRequired(),
		jwt.WithAudience(a.audience),
		jwt.WithIssuer(a.audience),
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
	)
}
