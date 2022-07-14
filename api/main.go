package main

import (
	"api/controller"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)

	switch r.Method {
	case "GET":
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
	case "POST":
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

func main() {
	http.HandleFunc("/users", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
