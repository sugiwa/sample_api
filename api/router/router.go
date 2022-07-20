package router

import (
	"api/controller"
	"fmt"
	"net/http"
)

func HandleRequest(w http.ResponseWriter, r *http.Request, param string) {
	switch param {
	case "login":
		if r.Method == http.MethodPost {
			controller.Login(w, r)
		}
	case "register":
		fmt.Println("REGISTER")
	}
}
