package jwt

import (
	"time"

	jwtlib "github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwtlib.RegisteredClaims
}

func GenerateToken(
	secret string,
	userID string,
	email string,
	role string,
) (string, error) {

	claims := Claims{
		UserID: userID,
		Email:  email,
		Role:   role,
		RegisteredClaims: jwtlib.RegisteredClaims{
			ExpiresAt: jwtlib.NewNumericDate(
				time.Now().Add(24 * time.Hour),
			),
			IssuedAt: jwtlib.NewNumericDate(
				time.Now(),
			),
		},
	}

	token := jwtlib.NewWithClaims(
		jwtlib.SigningMethodHS256,
		claims,
	)

	return token.SignedString(
		[]byte(secret),
	)
}

func ValidateToken(
	tokenString string,
	secret string,
) (*Claims, error) {

	token, err := jwtlib.ParseWithClaims(
		tokenString,
		&Claims{},
		func(token *jwtlib.Token) (interface{}, error) {
			return []byte(secret), nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)

	if !ok || !token.Valid {
		return nil, jwtlib.ErrTokenInvalidClaims
	}

	return claims, nil
}
