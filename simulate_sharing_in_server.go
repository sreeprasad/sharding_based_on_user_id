package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
)

const (
	db1DSN = "host=localhost port=5435 user=user4 password=password4 dbname=mydatabase4 sslmode=disable"
	db2DSN = "host=localhost port=5433 user=user5 password=password5 dbname=mydatabase5 sslmode=disable"
)

var db1 *sql.DB
var db2 *sql.DB

func initDB() {
	var err error
	db1, err = sql.Open("postgres", db1DSN)
	if err != nil {
		log.Fatalf("Could not connect to db1: %v", err)
	}

	db2, err = sql.Open("postgres", db2DSN)
	if err != nil {
		log.Fatalf("Could not connect to db2: %v", err)
	}
}

func executeHandler(w http.ResponseWriter, r *http.Request) {
	userIDParam := r.URL.Query().Get("userid")
	if userIDParam == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		http.Error(w, "Invalid User ID", http.StatusBadRequest)
		return
	}

	var dbID int
	var currentTime string

	if userID%2 == 0 {
		dbID = 1
		err = db1.QueryRow("SELECT NOW()").Scan(&currentTime)
	} else {
		dbID = 2
		err = db2.QueryRow("SELECT NOW()").Scan(&currentTime)
	}
	if err != nil {
		log.Fatal("Error querying the current time: ", err)
		http.Error(w, "Failed to execute database operation", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Current Time: %s using database ID %d", currentTime, dbID)
}

func main() {
	initDB()
	defer db1.Close()
	defer db2.Close()

	http.HandleFunc("/execute", executeHandler)

	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
