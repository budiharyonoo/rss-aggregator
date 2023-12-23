package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers *http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("no auth info found")
	}

	authHeaderValues := strings.Split(authHeader, " ")
	if len(authHeaderValues) < 2 || authHeaderValues[0] != "ApiKey" || authHeaderValues[1] == "" {
		return "", errors.New("malformed auth")
	}

	return authHeaderValues[1], nil
}
