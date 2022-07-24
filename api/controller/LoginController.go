package controller

import (
	"api/model"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) (err error) {
	var user model.User
	var requestUser model.User
	dec := json.NewDecoder(r.Body)
	err = dec.Decode(&requestUser)

	if err := DB.QueryRow(`SELECT * FROM users WHERE email = $1`, requestUser.Email).Scan(&user.Id, &user.Name, &user.Email, &user.Password); err != nil {
		w.WriteHeader(400)
		log.Println(err)
	}

	if requestUser.Password == user.Password {
		fmt.Println("login success")
		sessionID := createSessionID()
		fmt.Println(sessionID)
		cookie := &http.Cookie{
			Name: "session_id",
			Value: sessionID,
			MaxAge: 60,
		}

		http.SetCookie(w, cookie)

	} else {
		fmt.Println("login failed")
		w.WriteHeader(401)
	}

	return
}

func createSessionID() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		log.Fatal(err)
	}

	return base64.URLEncoding.EncodeToString(b)
}
