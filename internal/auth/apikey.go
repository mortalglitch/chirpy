package auth

import (
	"fmt"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	requestToken := headers.Get("Authorization")
	if requestToken == "" {
		return "", fmt.Errorf("Authorization header does not meet requirements")
	}
	trimmedKey := strings.TrimPrefix(requestToken, "ApiKey ")
	return trimmedKey, nil
}
