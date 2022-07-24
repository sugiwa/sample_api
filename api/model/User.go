package model

import (
	"api/util"
	"database/sql"
	"log"
	"net/http"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var DB *sql.DB

func init() {
	DB = util.GetDBConnector()
}

func CurrentUser(r *http.Request) User {
	uid := GetCurrentUserId(r)

	var user User
	err := DB.QueryRow(`SELECT * FROM users WHERE id = S1`, uid).Scan(&user.Id, &user.Email, &user.Name)

	if err != nil {
		log.Println(err)
	}

	return user
}
