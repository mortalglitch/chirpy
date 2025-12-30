package auth

import (
	"fmt"
	"net/http"
	"testing"
	"time"
	//"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func TestMakeJWT(t *testing.T) {
	userID, err := uuid.Parse("93e2d633-b9ab-4088-a694-37dbc58735fa")
	if err != nil {
		t.Errorf("Invalid UUID %v", err)
	}

	minutes, _ := time.ParseDuration("3m")

	secretString, err := MakeJWT(userID, "peanut", minutes)
	if err != nil {
		t.Errorf("failed to generate secret string: %v", err)
	}
	fmt.Print(secretString)

	testUser, err := ValidateJWT(secretString, "peanut")
	if err != nil {
		t.Errorf("failed to validate secret string: %v", err)
	}

	if testUser != userID {
		t.Error("User ID doesn't match after validation.")
	}

}

func TestGetBearerToken(t *testing.T) {
	headers := http.Header{
		"Authorization": []string{"Bearer my-jwt-token-123"},
	}
	
	returnedString, err := GetBearerToken(headers)
	if err != nil {
		t.Errorf("Error on processing token: %v", err)
	}

	if returnedString != "my-jwt-token-123" {
		t.Errorf("ReturnedString did not match: %s", returnedString)
	}
}


