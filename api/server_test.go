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

func TestSubscribeUser(t *testing.T) {
	//user should be subscribed on the correct topics, and only once in each of them
	createTopics()

	//invalid level - should not be subscribed
	t.Run("invalid input", func(t *testing.T) {
		id := 111
		level := "X"

		subscriptions := subscribe(id, level)

		if subscriptions > 0 {
			t.Errorf("Invalid level user was subscribed")
		}
	})

	//valid id and level - should subscribe to a certain number of topics
	t.Run("valid level - number of subscriptions", func(t *testing.T) {
		id := 1
		level := "A"

		subscriptions := subscribe(id, level)

		if subscriptions != 4 {
			t.Errorf("Invalid level user was subscribed: level: %s, subscriptions: %d", level, subscriptions)
		}

		id = 2
		level = "B"

		subscriptions = subscribe(id, level)

		if subscriptions != 3 {
			t.Errorf("Invalid level user was subscribed: level: %s, subscriptions: %d", level, subscriptions)
		}

		id = 3
		level = "C"

		subscriptions = subscribe(id, level)

		if subscriptions != 2 {
			t.Errorf("Invalid level user was subscribed: level: %s, subscriptions: %d", level, subscriptions)
		}

		id = 4
		level = "D"

		subscriptions = subscribe(id, level)

		if subscriptions != 1 {
			t.Errorf("Invalid level user was subscribed: level: %s, subscriptions: %d", level, subscriptions)
		}

	})

	//user should not be subcribed twice on same topic
	t.Run("user should only be subscribed once on the same topic", func(t *testing.T) {
		var newSub Subscriber
		newSub.UserId = 10

		first := subscribeOnce(newSub, findTopic("A"))
		second := subscribeOnce(newSub, findTopic("A"))

		if first == false || second == true {
			t.Errorf("User was not subscribed once on a topic")
		}
	})

}
