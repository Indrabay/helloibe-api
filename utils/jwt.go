package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT interface {
	CreateToken(JWTClaim) (string, error)
	ValidateToken(token string) (JWTClaim, error)
}

type JWTStruct struct {
	SingingKey string
}

type JWTClaim struct {
	Username string
	Name     string
	Role     int
	Level    int
	Stores   []int64
	jwt.RegisteredClaims
}

func NewJWT(singingKey string) *JWTStruct {
	return &JWTStruct{SingingKey: singingKey}
}

func (j *JWTStruct) CreateToken(claim JWTClaim) (string, error) {
	claims := JWTClaim{
		claim.Username,
		claim.Name,
		claim.Role,
		claim.Level,
		claim.Stores,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Hour)),
			Issuer:    "helloibe.com-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, err := token.SignedString([]byte(Config.SigningKey))
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func (j *JWTStruct) ValidateToken(token string) (JWTClaim, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &JWTClaim{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(Config.SigningKey), nil
	})
	if err != nil {
		return JWTClaim{}, ErrUserNotAuthorized
	}

	if claims, ok := parsedToken.Claims.(*JWTClaim); parsedToken.Valid && ok {
		return *claims, nil
	}

	return JWTClaim{}, ErrTokenNotValid
}
