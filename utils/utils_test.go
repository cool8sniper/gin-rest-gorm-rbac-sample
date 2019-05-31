package utils

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	if i, e := HashPassword("abcdef"); len(i) <= 30 || e != nil {
		t.Errorf("hashPassword wrong")
	}
}

func TestCheckPasswordHash(t *testing.T) {
	password := "abcdefg"
	hash, _ := HashPassword(password)

	if CheckPasswordHash(password, hash) {
		t.Log("password is good")
	}

	if !CheckPasswordHash("wrongpassword", hash) {
		t.Log("password is good")
	}
}
