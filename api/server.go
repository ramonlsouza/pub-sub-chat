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

func auth(res http.ResponseWriter, req *http.Request) {
	//enable CORS
	res.Header().Set("Access-Control-Allow-Origin", "*")

	if req.Method != "POST" {
		fmt.Fprintf(res, "Authorization should be made with a POST request!\n")
	} else {
		// read and parse json file
		var userlist Users
		var validUserData User
		var apiReturn ApiReturn

		file, openErr := os.Open("users.json")

		if openErr != nil {
			fmt.Println(openErr)
			return
		}

		bytes, readErr := ioutil.ReadAll(file)

		if readErr != nil {
			fmt.Println(readErr)
			return
		}

		//push parsed users to userlist
		json.Unmarshal(bytes, &userlist)

		defer file.Close()

		//search for valid user
		req.ParseForm()

		user := req.FormValue("username")
		pass := req.FormValue("password")

		validUser := false

		for i := 0; i < len(userlist.Users); i++ {
			if userlist.Users[i].Username == user && userlist.Users[i].Password == pass {
				validUser = true
				validUserData = userlist.Users[i]
			}
		}

		res.Header().Set("Content-Type", "application/json")

		if validUser == true {
			apiReturn.Error = false
			apiReturn.Message = ""
			apiReturn.UserData = validUserData

			apiReturn.Token = generateToken(validUserData.Id, validUserData.Level)

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
