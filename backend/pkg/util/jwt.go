package util

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"

)

var jwtSecret []byte

type Claims struct {
	Stid string `json:"stid"`
	Privilege int `json:"privilege"`
	jwt.StandardClaims
}

func GenerateToken(stid string, privilege int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		stid,
		privilege,
		jwt.StandardClaims {
			ExpiresAt : expireTime.Unix(),
			Issuer : "ts-system",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}