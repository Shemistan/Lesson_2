package controller

import (
	"fmt"
	"github.com/Shemistan/Lesson_2/model"
)

func SignIn(logData model.UserLogin, db *DB) model.User {
	var user model.User

	if logData.Login == "" || logData.Password == "" {
		fmt.Println("Пустое поле не допустимо")
	}

	for _, userData := range db.Users {
		if userData.Login == logData.Login && userData.Password == logData.Password {
			user = userData
		}
	}

	return user
}

func SignUp(newUser model.User, db *DB) model.User {

	if newUser.Balance == 0 || newUser.FullName == "" || newUser.PhoneNumber == "" || newUser.Login == "" || newUser.Password == "" {
		fmt.Println("Пустое поле не допустимо")
	}

	db.Users = append(db.Users, newUser)

	return newUser
}
