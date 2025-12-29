package auth

import (
	"testing"

	"github.com/alexedwards/argon2id"
)

func TestHashPassword(t *testing.T) {
	password := "p@ssword"

	hash, err := HashPassword(password)
	if err != nil {
		t.Errorf("Password failed to hash %v", err)
	}

	match, err := argon2id.ComparePasswordAndHash(password, hash)
	if err != nil {
		t.Fatalf("argon2id.VerifyPassword faild: %v", err)
	}

	if !match {
		t.Errorf("Password verification failed. The generated hash is invalid.")
	}

	if hash == password {
		t.Error("Hash output is the same as the input password (it was not hashed)")
	}
}
