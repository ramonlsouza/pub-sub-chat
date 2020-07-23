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

		_, _, isValid := parseToken(token)

		if isValid != false {
			t.Errorf("Valid return for empty token")
		}
	})

	//invalid input - should return error
	t.Run("invalid input", func(t *testing.T) {
		token := "aaaaaa"

		_, _, isValid := parseToken(token)

		if isValid != false {
			t.Errorf("Valid return for invalid token")
		}
	})

	//valid input - should return token data
	t.Run("valid input", func(t *testing.T) {
		token := generateToken(1, "A")

		_, _, isValid := parseToken(token)

		if isValid != true {
			t.Errorf("Failed auth with valid token")
		}
	})
}

func TestSendMessage(t *testing.T) {
	/* a valid input for an attempt to send a message should contain:
	1. a valid token
	2. a message string
	*/

	//invalid input - should return error
	t.Run("empty message and invalid token", func(t *testing.T) {
		token := "aaaaaa"
		message := ""

		isSent := sendMessage(token, message)

		if isSent != false {
			t.Errorf("Valid return for empty message and invalid token")
		}
	})

	//invalid input - should return error
	t.Run("message and token, but invalid token", func(t *testing.T) {
		token := "aaaaaa"
		message := "my test message"

		isSent := sendMessage(token, message)

		if isSent != false {
			t.Errorf("Valid return for invalid token")
		}
	})

	//empty message - should return error
	t.Run("empty message", func(t *testing.T) {
		token := generateToken(1, "A")
		message := ""

		isSent := sendMessage(token, message)

		if isSent != false {
			t.Errorf("Valid return for valid token and empty message")
		}
	})

	//valid input - should return ok
	t.Run("valid input", func(t *testing.T) {
		token := generateToken(1, "A")
		message := "my test message"

		isSent := sendMessage(token, message)

		if isSent != true {
			t.Errorf("Failed to send message with valid token")
		}
	})
}
