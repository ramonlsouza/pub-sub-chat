package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"net/http"
	"os"
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

type ApiReturn struct {
	Error    bool
	UserData User
	Message  string
	Token    string
}

func generateToken(userId int, userLevel string) string {
	//on a real project, the secret key should be in a separated file
	secret := "not-so-secret-key"

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["user_level"] = userLevel
	atClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, _ := at.SignedString([]byte(secret))
	return token
}

func parseToken(token string) (string, bool) {
	//TODO
	return "", false
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

func auth(res http.ResponseWriter, req *http.Request) {
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

		var apiReturn ApiReturn

		if validUser == true {
			apiReturn.Error = false
			apiReturn.Message = ""
			apiReturn.UserData = userData

			apiReturn.Token = generateToken(userData.Id, userData.Level)

			jsonReturn, _ := json.Marshal(apiReturn)
			fmt.Fprintf(res, string(jsonReturn))
		} else {
			apiReturn.Error = true
			apiReturn.Message = "invalid credentials!"

			jsonReturn, _ := json.Marshal(apiReturn)
			fmt.Fprintf(res, string(jsonReturn))
		}
	}
}

func main() {
	http.HandleFunc("/auth", auth)

	http.ListenAndServe(":8000", nil)
}
