package controller

import (
	"api/model"
	"api/util"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var DB *sql.DB

func init() {
	DB = util.GetDBConnector()
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

	fmt.Println(users)

	return
}

func GetUser(id int, r *http.Request) (user model.User, err error) {
	fmt.Println("start GetUser")

	dec := json.NewDecoder(r.Body)
	err = dec.Decode(&user)

	if err := DB.QueryRow("SELECT id, name, email, password FROM users where id = $1", id).Scan(&user.Id, &user.Name, &user.Email, &user.Password); err != nil {
		fmt.Println(err)
	}

	fmt.Println(user)

	return
}

func InsertUser(w http.ResponseWriter, r *http.Request) (req model.User, err error) {
	fmt.Println("start InsertUser")

	dec := json.NewDecoder(r.Body)
	err = dec.Decode(&req)

	res, err := DB.Exec(`INSERT INTO users (name, email, password) VALUES ($1, $2, $3)`, req.Name, req.Email, req.Password)
	fmt.Println("request fin")
	if err != nil {
		w.WriteHeader(400)
		log.Println(err)
	}

	fmt.Println(res)

	return
}

func DeleteUser(id int, r *http.Request) (res sql.Result, err error) {
	fmt.Println("start deleteUser")

	// var req model.User
	// dec := json.NewDecoder(r.Body)
	// err = dec.Decode(&req)

	res, err = DB.Exec("DELETE from users where id = $1", id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)

	return
}
