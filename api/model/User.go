package model

import "fmt"

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Say() {
	fmt.Println("test")
}