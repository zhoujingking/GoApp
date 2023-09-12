package authorization

import "github.com/golang-jwt/jwt"

type UserClaim struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	jwt.StandardClaims
}
