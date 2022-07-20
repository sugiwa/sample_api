package util

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func init() {
	CreateDB()
}

func CreateDB() {
	var err error
	user := os.Getenv("POSTGRES_USER")
	db := os.Getenv("POSTGRES_DB")
	password := os.Getenv("POSTGRES_PASSWORD")
	config := fmt.Sprintf("host=postgres user=%s dbname=%s password=%s sslmode=disable", user, db, password)
	DB, err = sql.Open("postgres", config)
	if err != nil {
		panic(err)
	}

	for {
		err = DB.Ping()
		if err == nil {
			fmt.Println("Connection to DB Success")
			break
		}
		fmt.Println("Connection to DB Failed", err)
		time.Sleep(3 * time.Millisecond)
	}
}

func GetDBConnector() *sql.DB {
	for {
		if DB != nil {
			break
		}
		time.Sleep(3 * time.Millisecond)
	}
	return DB
}