package models

import (
	"database/sql"
	"sagar/entities"
)

type UserModel struct {
	Db *sql.DB
}

func (userModel UserModel) FindAll() ([]entities.User, error) {
	rows, err := userModel.Db.Query("select * from [dbo].[Users]")

	if err != nil {
		return nil, err
	} else {
		var users []entities.User = make([]entities.User, 0)
		for rows.Next() {
			var username string
			var password string
			rows.Scan(&username, &password)
			user := entities.User{Username: username, Password: password}
			users = append(users, user)
		}
		return users, err
	}
}

func (userModel UserModel) Create(user entities.User) (int64, error) {
	result, err := userModel.Db.Exec("INSERT INTO [dbo].[USERS] VALUES ('" + user.Username + "','" + user.Password + "')")

	if err != nil {
		return 0, err
	} else {
		id, _ := result.LastInsertId()
		return id, err
	}
}
