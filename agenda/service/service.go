package service

import "github.com/whichxjy/Service-Computing-Homework/agenda/model"

func Register(username string, password string, email string) bool {
	return model.AddUser(username, password, email)
}

func Login(username string, password string) bool {
	return model.FindUser(username, password)
}
