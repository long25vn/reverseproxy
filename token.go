package main

import (
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

type UserClaim struct {
	Username string
	Role    int32
	jwt.StandardClaims
}

func verifyToken(jwtToken string) (*UserClaim, error) {
	token, err := jwt.ParseWithClaims(jwtToken, &UserClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})
	if err == nil {
		if claims, ok := token.Claims.(*UserClaim); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, err
}

func getJwtToken(header http.Header) string {
	const bearerSchema = "Bearer "
	authHeader := header["Authorization"]
	if len(authHeader) == 0 || authHeader[0] == "" {
		return ""
	}
	if !strings.HasPrefix(authHeader[0], bearerSchema) {
		return ""
	}
	return authHeader[0][len(bearerSchema):]
}
