package model

import (
	"os"
	"encoding/json"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

var userFileName = "data/user.json"

func AddUser(username string, password string, email string) bool {
	var users []User
	if err := LoadJson(userFileName, &users); err != nil {
		panic("Fail to load user data")
	}
	for _, user := range users {
		if user.Username == username {
			return false
		}
	}
	users = append(users, User{ username, password, email })
	if err := SaveJson(userFileName, &users); err != nil {
		panic("Fail to save user data")
	}
	return true
}

func FindUser(username string, password string) bool {
	var users []User
	if err := LoadJson(userFileName, &users); err != nil {
		panic("Fail to load user data")
	}
	for _, user := range users {
		if user.Username == username && user.Password == password {
			return true
		}
	}
	return false;
}

func LoadJson(filename string, v interface{}) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	// load json to v
	return json.NewDecoder(f).Decode(v)
}

func SaveJson(filename string, v interface{}) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	// save json to file
	return json.NewEncoder(f).Encode(v)
}
