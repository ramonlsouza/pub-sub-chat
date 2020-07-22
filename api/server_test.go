package main

import "testing"

func TestAuth(t *testing.T) {
	//empty input - should return error
	t.Run("empty input", func(t *testing.T) {
		_, validUser := checkUser("", "")

		if validUser != false {
			t.Errorf("Valid return for auth with empty credentials")
		}

	})

	//invalid input - should return error
	t.Run("invalid input", func(t *testing.T) {
		_, validUser := checkUser("123", "123")

		if validUser != false {
			t.Errorf("Valid return for auth with invalid credentials")
		}
	})

	//valid input - should return user data
	t.Run("valid input", func(t *testing.T) {
		_, validUser := checkUser("A", "A123")

		if validUser != true {
			t.Errorf("Failed auth with valid credentials")
		}
	})
}

func TestToken(t *testing.T) {
	//empty input - should return error
	t.Run("empty input", func(t *testing.T) {
		token := ""

		_, validToken := parseToken(token)

		if validToken != false {
			t.Errorf("Valid return for empty token")
		}
	})

	//invalid input - should return error
	t.Run("invalid input", func(t *testing.T) {
		token := "aaaaaa"

		_, validToken := parseToken(token)

		if validToken != false {
			t.Errorf("Valid return for invalid token")
		}
	})

	//valid input - should return token data
	t.Run("valid input", func(t *testing.T) {
		token := generateToken(1, "A")

		_, validToken := parseToken(token)

		if validToken != true {
			t.Errorf("Failed auth with valid token")
		}
	})
}
