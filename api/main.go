package main

import (
	"api/model"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var users = []model.User{
	{"test1", "test1@gmail.com", "test"},
	{"test2", "test2@gmail.com", "test"},
	{"test3", "test3@gmail.com", "test"},
}

func handler(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(users); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String())

	_, err := fmt.Fprint(w, buf.String())
	if err != nil {
		return
	}

}

func main() {
	http.HandleFunc("/users", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
