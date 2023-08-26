package authorization

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type UserClaim struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	jwt.StandardClaims
}

const secret = "godking-oceansky"

func GenerateToken(name string, id string) (string, error) {
	const tokenDuration = time.Hour * 2

	claims := UserClaim{
		Id:   id,
		Name: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenDuration).Unix(),
			Issuer:    "godking",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenSigned, err := token.SignedString([]byte(secret))
	return tokenSigned, err
}

func VerifyToken(tokenString string) (*UserClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if claims, ok := token.Claims.(*UserClaim); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
