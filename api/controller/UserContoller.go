package controller

import (
	"api/model"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("postgres", "host=postgres user=hoge dbname=db password=root sslmode=disable")
	if err != nil {
		panic(err)
	}

	for {
		err = DB.Ping()
		if err == nil {
			fmt.Println("Connection to DB Success")
			break
		}
		fmt.Println("Connection to DB Failed", err)
		time.Sleep(3 * time.Millisecond)
	}
}

func GetUsers() (users []model.User, err error) {
	fmt.Println("start GetUsers")

	rows, err := DB.Query("SELECT id, name, email, password FROM users")

	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		var user model.User
		rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
		users = append(users, user)
	}

	return
}
