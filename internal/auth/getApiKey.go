package auth

import (
	"net/http"
	"strings"
)

func GetApiKey(header http.Header) (string, error) {
	authHeader := header.Get("Authorization")
	if authHeader == "" {
		return "", http.ErrNoCookie
	}

	authKey := strings.Split(authHeader, " ")
	if len(authKey) != 2 || authKey[0] != "Bearer" {
		return "", http.ErrNoCookie
	}

	return authKey[1], nil
}
