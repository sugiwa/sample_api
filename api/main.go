package main

import (
	"api/controller"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)

	switch r.Method {
	case http.MethodGet:
		req, _ := controller.GetUsers()
		fmt.Println(req)
		if err := enc.Encode(req); err != nil {
			log.Fatal(err)
		}
		fmt.Println(buf.String())

		_, err := fmt.Fprint(w, buf.String())
		if err != nil {
			return
		}
	case http.MethodPost:
		req, _ := controller.InsertUser(w, r)
		fmt.Println(req)
		if err := enc.Encode(req); err != nil {
			log.Fatal(err)
		}
		fmt.Println(buf.String())

		_, err := fmt.Fprint(w, buf.String())
		if err != nil {
			return
		}
	}

}

func handler2(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)

	fmt.Println(r.URL.Path)
	sub := strings.TrimPrefix(r.URL.Path, "/users")
	fmt.Println(sub)
	_, param := filepath.Split(sub)
	if param != "" {
		fmt.Println(param)
	}

	id, _ := strconv.Atoi(param)

	switch r.Method {
	case http.MethodGet:
		req, _ := controller.GetUser(id, r)
		fmt.Println(req)
		if err := enc.Encode(req); err != nil {
			log.Fatal(err)
		}

		_, err := fmt.Fprint(w, buf.String())
		if err != nil {
			return
		}
	case http.MethodDelete:
		req, _ := controller.DeleteUser(id, r)
		fmt.Println(req)
		if err := enc.Encode(req); err != nil {
			log.Fatal(err)
		}

		_, err := fmt.Fprint(w, buf.String())
		if err != nil {
			return
		}

	}
}

func main() {
	http.HandleFunc("/users", handler)
	http.HandleFunc("/users/", handler2)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
