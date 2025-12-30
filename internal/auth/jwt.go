package auth

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type CustomClaims struct {
	jwt.RegisteredClaims
}

func MakeJWT(userID uuid.UUID, tokenSecret string, expiresIn time.Duration) (string, error) {
	claims := CustomClaims {
		jwt.RegisteredClaims{
			Issuer:    "chirpy",
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(expiresIn)),
			Subject:   userID.String(),
		},
	}

	secretArray := []byte(tokenSecret)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(secretArray)

	return ss, err
}

func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error) {
	var user uuid.UUID

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(tokenSecret), nil
	})
	if err != nil {
		log.Printf("Error processing JWT: %v", err)
		return user, err
	} 
	
	if claims, ok := token.Claims.(*CustomClaims); ok {
		currentUserID, err := uuid.Parse(claims.Subject)
		if err != nil {
			log.Printf("unable to process uuid")
			return user, nil
		}
		user = currentUserID
	} else {
		log.Printf("unknown claims type, cannot proceed")
	}
	return user, nil
}

func GetBearerToken(headers http.Header) (string, error) {
	requestToken := headers.Get("Authorization")
	if requestToken == "" {
		return "", fmt.Errorf("Authorization header does not meet requirements")
	}
	trimmedToken := strings.TrimPrefix(requestToken, "Bearer ")
	return trimmedToken, nil
}
