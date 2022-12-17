package services

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtService interface {
	GenerateToken(administratorID uint, now time.Time) (string, error)
	ParseToken(signedString string) (*Auth, error)
}

type jwtService struct{}

func NewJwtService() JwtService {
	return &jwtService{}
}

type Auth struct {
	Uid    uint  `json:"uid"`
	Iat    int64 `json:"iat"`
	Expiry int64 `json:"expiry"`
}

func (service *jwtService) GenerateToken(userId uint, now time.Time) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid":    userId,
		"iat":    now.Unix(),
		"expiry": now.Add(24 * time.Hour).Unix(),
	})

	return token.SignedString([]byte(os.Getenv("TOKEN_KEY")))
}

func validToken(signedString string) (*jwt.Token, error) {
	token, err := jwt.Parse(signedString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("TOKEN_KEY")), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func (service *jwtService) ParseToken(signedString string) (*Auth, error) {
	token, err := validToken(signedString)

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, fmt.Errorf("%s is expired", signedString)
			} else {
				return nil, fmt.Errorf("%s is invalid", signedString)
			}
		} else {
			return nil, fmt.Errorf("%s is invalid", signedString)
		}
	}

	if token == nil {
		return nil, fmt.Errorf("not found token in %s:", signedString)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("not found claims in %s", signedString)
	}

	userId, ok := claims["uid"].(float64)
	if !ok {
		return nil, fmt.Errorf("not found uid in %s", signedString)
	}

	iat, ok := claims["iat"].(float64)
	if !ok {
		return nil, fmt.Errorf("not found iat in %s", signedString)
	}
	expiry, ok := claims["expiry"].(float64)
	if !ok {
		return nil, fmt.Errorf("not found expiry in %s", signedString)
	}

	return &Auth{
		Uid:    uint(userId),
		Iat:    int64(iat),
		Expiry: int64(expiry),
	}, nil
}
