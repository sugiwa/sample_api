package model

import (
	"fmt"
	"log"
	"net/http"
)

type SessionInfo struct {
	Id        int    `json:"id"`
	SessionId string `json:"session_id"`
	UserId    int    `json:"user_id"`
}

func GetCurrentUserId(r *http.Request) int {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(cookie)

	var sid SessionInfo
	err = DB.QueryRow(`SELECT user_id FROM session_ids WHERE session_id = S1`, cookie.Value).Scan(&sid.UserId)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(sid.UserId)

	return sid.UserId
}
