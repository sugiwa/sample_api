package main

import (
	"api/controller"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)

	switch r.Method {
	case http.MethodGet:
		req, _ := controller.GetUsers()
		if err := enc.Encode(req); err != nil {
			log.Fatal(err)
		}

		_, err := fmt.Fprint(w, buf.String())
		if err != nil {
			return
		}
	case http.MethodPost:
		req, _ := controller.InsertUser(w, r)
		if err := enc.Encode(req); err != nil {
			log.Fatal(err)
		}

		_, err := fmt.Fprint(w, buf.String())
		if err != nil {
			return
		}
	case http.MethodPut:
		w.WriteHeader(404)
	case http.MethodDelete:
		w.WriteHeader(404)
	}

}

func handler2(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)

	sub := strings.Trim(r.URL.Path, "/")
	params := strings.Split(sub, "/")

	if !isNumber(params[1]) {
		return
	}

	id, _ := strconv.Atoi(params[1])

	switch r.Method {
	case http.MethodGet:
		req, _ := controller.GetUser(id, r)
		if err := enc.Encode(req); err != nil {
			log.Fatal(err)
		}

		_, err := fmt.Fprint(w, buf.String())
		if err != nil {
			return
		}
	case http.MethodPost:
		w.WriteHeader(404)
	case http.MethodPut:
		w.WriteHeader(404)
	case http.MethodDelete:
		req, _ := controller.DeleteUser(id, r)
		if err := enc.Encode(req); err != nil {
			log.Fatal(err)
		}

		_, err := fmt.Fprint(w, buf.String())
		if err != nil {
			return
		}

	}
}

func isNumber(str string) bool {
	reg := `^[0-9]\d*$`
	return regexp.MustCompile(reg).Match([]byte(str))
}

func main() {
	http.HandleFunc("/users", handler)
	http.HandleFunc("/users/", handler2)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
