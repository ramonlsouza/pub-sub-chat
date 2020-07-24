package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

//STRUCTS
type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Level    string `json:"level"`
	Password string `json:"password"`
}

type AuthReturn struct {
	Error    bool
	UserData User
	Message  string
	Token    string
}

type ApiReturn struct {
	Error   bool
	Message string
}

type MessagesReturn struct {
	Error        bool
	Message      string
	UserMessages []Message
}

type Topic struct {
	Name        string
	Subscribers []Subscriber
}

type Subscriber struct {
	UserId int
}

type UserMessages struct {
	UserId   int
	Messages []Message
}

//TODO: add message date and time?
type Message struct {
	SenderId    int
	SenderName  string
	MessageText string
}

//INSTANCES
var topics []Topic
var messages []UserMessages
var userlist []User

//FUNCTIONS
func createTopics() {
	var newTopics = []string{"A", "B", "C", "D"}

	for _, v := range newTopics {
		var topic Topic
		topic.Name = v
		topics = append(topics, topic)
	}
}

func createUsers() {
	var newUsers = []string{"A", "B", "C", "D"}

	for i, v := range newUsers {
		var user User
		user.Id = i + 1
		user.Username = v
		user.Level = v
		user.Password = "123"
		userlist = append(userlist, user)
	}
}

func findUser(userId int) (index int) {
	for i := 0; i < len(userlist); i++ {
		if userlist[i].Id == userId {
			return i
		}
	}
	return -1
}

func findTopic(name string) (index int) {
	for i := 0; i < len(topics); i++ {
		if topics[i].Name == name {
			return i
		}
	}
	return -1
}

func findUserMessages(userId int) (index int) {
	index = len(messages)

	for i := 0; i < len(messages); i++ {
		if messages[i].UserId == userId {
			return i
		}
	}
	var emptyUserMessages UserMessages
	emptyUserMessages.UserId = userId

	messages = append(messages, emptyUserMessages)
	return index
}

func subscribeOnce(sub Subscriber, index int) bool {
	exists := false

	for i := 0; i < len(topics[index].Subscribers); i++ {
		if topics[index].Subscribers[i].UserId == sub.UserId {
			exists = true
		}
	}

	if exists == true {
		return false
	} else {
		topics[index].Subscribers = append(topics[index].Subscribers, sub)
		return true
	}
}

func subscribe(userId int, userLevel string) int {
	var newSub Subscriber
	newSub.UserId = userId

	subscriptions := 0

	switch userLevel {
	case "A":
		if subscribeOnce(newSub, findTopic("A")) == true {
			subscriptions++
		}
		fallthrough
	case "B":
		if subscribeOnce(newSub, findTopic("B")) == true {
			subscriptions++
		}
		fallthrough
	case "C":
		if subscribeOnce(newSub, findTopic("C")) == true {
			subscriptions++
		}
		fallthrough
	case "D":
		if subscribeOnce(newSub, findTopic("D")) == true {
			subscriptions++
		}
	}

	return subscriptions
}

func generateToken(userId int, userLevel string) string {
	//on a real project, the secret key should be in a separated file
	secret := "not-so-secret-key"

	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = userId
	atClaims["user_level"] = userLevel
	atClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, _ := at.SignedString([]byte(secret))
	return token
}

func parseToken(tokenString string) (userId int, userLevel string, isValid bool) {
	//on a real project, the secret key should be in a separated file
	secret := "not-so-secret-key"

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err == nil {
		claims := token.Claims.(jwt.MapClaims)

		userId := int(claims["user_id"].(float64))
		userLevel := claims["user_level"].(string)

		//TODO: check token expiration
		return userId, userLevel, true
	} else {
		return 0, "", false
	}

}

func checkUser(username string, password string) (User, bool) {
	var userData User

	for i := 0; i < len(userlist); i++ {
		if userlist[i].Username == username && userlist[i].Password == password {
			userData = userlist[i]
			return userData, true
		}
	}
	return userData, false
}

func validateMessage(token string, message string) (int, User, bool) {
	userId, userLevel, isValid := parseToken(token)
	topicIndex := findTopic(userLevel)
	userIndex := findUser(userId)

	if isValid == true && message != "" && topicIndex >= 0 && userIndex >= 0 {
		userData := userlist[userIndex]
		return topicIndex, userData, true
	}

	var emptyUser User
	return 0, emptyUser, false
}

func sendMessage(topicIndex int, userData User, message string) {
	//call this only after validateMessage returns true!
	var newMessage Message
	newMessage.MessageText = message
	newMessage.SenderId = userData.Id
	newMessage.SenderName = userData.Username

	for i := 0; i < len(topics[topicIndex].Subscribers); i++ {
		userIndex := findUserMessages(topics[topicIndex].Subscribers[i].UserId)
		messages[userIndex].Messages = append(messages[userIndex].Messages, newMessage)
	}
}

func getUserMessages(token string) []Message {
	userId, _, _ := parseToken(token)

	userIndex := findUserMessages(userId)
	userMessages := messages[userIndex].Messages

	return userMessages
}

//ROUTES
func authRoute(res http.ResponseWriter, req *http.Request) {
	//enable CORS
	res.Header().Set("Access-Control-Allow-Origin", "*")

	if req.Method != "POST" {
		fmt.Fprintf(res, "Authorization should be made with a POST request!\n")
	} else {
		//search for valid user
		req.ParseForm()

		user := req.FormValue("username")
		pass := req.FormValue("password")

		res.Header().Set("Content-Type", "application/json")

		userData, validUser := checkUser(user, pass)

		var authReturn AuthReturn

		if validUser == true {
			authReturn.Error = false
			authReturn.Message = ""
			authReturn.UserData = userData

			authReturn.Token = generateToken(userData.Id, userData.Level)

			//subscribe on topics
			go subscribe(userData.Id, userData.Level)

			jsonReturn, _ := json.Marshal(authReturn)
			fmt.Fprintf(res, string(jsonReturn))
		} else {
			authReturn.Error = true
			authReturn.Message = "invalid credentials!"

			jsonReturn, _ := json.Marshal(authReturn)
			fmt.Fprintf(res, string(jsonReturn))
		}
	}
}

func sendMessageRoute(res http.ResponseWriter, req *http.Request) {
	//enable CORS
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Headers", "Token,Content-Type")

	if req.Method != "POST" {
		fmt.Fprintf(res, "Authorization should be made with a POST request!\n")
	} else {
		var apiReturn ApiReturn

		req.ParseForm()
		message := req.FormValue("message")
		token := req.Header.Get("Token")

		//subscribe on topics
		userId, userLevel, _ := parseToken(token)
		go subscribe(userId, userLevel)

		res.Header().Set("Content-Type", "application/json")

		topicIndex, userData, isValid := validateMessage(token, message)

		if isValid == true {
			go sendMessage(topicIndex, userData, message)
			apiReturn.Error = false
			apiReturn.Message = "message sent!"
		} else {
			apiReturn.Error = true
			apiReturn.Message = "invalid data!"
		}

		jsonReturn, _ := json.Marshal(apiReturn)
		fmt.Fprintf(res, string(jsonReturn))
	}
}

func getMessagesRoute(res http.ResponseWriter, req *http.Request) {
	//enable CORS
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Headers", "Token,Content-Type")

	req.ParseForm()
	token := req.Header.Get("Token")

	//subscribe on topics
	userId, userLevel, _ := parseToken(token)
	go subscribe(userId, userLevel)

	res.Header().Set("Content-Type", "application/json")

	if token != "" {
		userMessages := getUserMessages(token)

		var messagesReturn MessagesReturn
		messagesReturn.Error = false
		messagesReturn.UserMessages = userMessages

		jsonReturn, _ := json.Marshal(messagesReturn)
		fmt.Fprintf(res, string(jsonReturn))
	} else {
		var apiReturn ApiReturn
		apiReturn.Error = true
		apiReturn.Message = "token not informed"

		jsonReturn, _ := json.Marshal(apiReturn)
		fmt.Fprintf(res, string(jsonReturn))
	}
}

func main() {
	createTopics()
	createUsers()

	http.HandleFunc("/auth", authRoute)
	http.HandleFunc("/send-message", sendMessageRoute)
	http.HandleFunc("/get-messages", getMessagesRoute)

	fmt.Println("server is running on port 8000")
	http.ListenAndServe(":8000", nil)
}
