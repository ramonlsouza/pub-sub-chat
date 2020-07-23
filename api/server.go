package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

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

type Topic struct {
	Name        string
	Subscribers []Subscriber
}

type Subscriber struct {
	UserId   int
	Messages []Message
}

//TODO: add message date and time?
type Message struct {
	SenderId    int
	MessageText string
	IsNew       bool
}

var topics []Topic

func createTopics() {
	var topicA Topic
	topicA.Name = "A"

	var topicB Topic
	topicB.Name = "B"

	var topicC Topic
	topicC.Name = "C"

	var topicD Topic
	topicD.Name = "D"

	topics = append(topics, topicA)
	topics = append(topics, topicB)
	topics = append(topics, topicC)
	topics = append(topics, topicD)
}

func findTopic(name string) (index int) {
	index = len(topics) - 1

	for i := 0; i < len(topics); i++ {
		if topics[i].Name == name {
			return i
		}
	}
	return -1
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
	// read and parse json file
	var userlist Users
	var userData User

	file, openErr := os.Open("users.json")

	if openErr != nil {
		fmt.Println(openErr)
		return userData, false
	}

	bytes, readErr := ioutil.ReadAll(file)

	if readErr != nil {
		fmt.Println(readErr)
		return userData, false
	}

	//push parsed users to userlist
	json.Unmarshal(bytes, &userlist)

	defer file.Close()

	for i := 0; i < len(userlist.Users); i++ {
		if userlist.Users[i].Username == username && userlist.Users[i].Password == password {
			userData = userlist.Users[i]
			return userData, true
		}
	}
	return userData, false
}

func sendMessage(token string, message string) bool {
	//TODO
	return false
}

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
			subscribe(userData.Id, userData.Level)

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
		tokenHeader := req.Header.Get("Token")

		splitToken := strings.Split(tokenHeader, " ")

		res.Header().Set("Content-Type", "application/json")

		isSent := sendMessage(splitToken[1], message)

		if isSent == true {
			apiReturn.Error = false
			apiReturn.Message = "message sent!"
		} else {
			apiReturn.Error = true
			apiReturn.Message = "error!"
		}

		jsonReturn, _ := json.Marshal(apiReturn)
		fmt.Fprintf(res, string(jsonReturn))
	}
}

func main() {
	createTopics()

	http.HandleFunc("/auth", authRoute)
	http.HandleFunc("/send-message", sendMessageRoute)

	http.ListenAndServe(":8000", nil)
}
