package main

import (
	"sagar/config"
	"sagar/entities"
	"sagar/models"
)

func main() {
	db, err := config.GetDB()

	if err != nil {
		panic(err)
	} else {
		usersModel := models.UserModel{Db: db}

		user := entities.User{Username: "Hanah", Password: "789"}

		v, _ := usersModel.Create(user)
		println("Id : ", v)
		users, dberr := usersModel.FindAll()
		if dberr != nil {
			panic(dberr)
		}
		for _, user := range users {
			println(user.Username)
		}
	}

}
