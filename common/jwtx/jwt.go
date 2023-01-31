package jwtx

import (
	"github.com/golang-jwt/jwt/v4"
)

func GetToken(secretKey string, iat, seconds, uid int64) (string, error) {
	claim := make(jwt.MapClaims)
	claim["exp"] = iat + seconds
	claim["iat"] = iat
	claim["uid"] = uid
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claim
	return token.SignedString(secretKey)
}
