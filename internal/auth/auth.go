package auth

import (
	"errors"
	"net/http"
	"strings"
)

// extracts the api key from the headers
// example:
// Authorization: Bearer {api_key}
func GetApiKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("authorization header not found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 || vals[0] != "ApiKey" {
		return "", errors.New("invalid authorization header format")
	}

	return vals[1], nil
}
