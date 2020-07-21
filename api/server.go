package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Username string `json:"username"`
	Level    string `json:"level"`
	Password string `json:"password"`
}

type ApiReturn struct {
	Error    bool
	UserData User
	Message  string
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
		var validUser = false
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
