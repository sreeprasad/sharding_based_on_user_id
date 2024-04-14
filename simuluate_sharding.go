package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	db1DSN = "host=localhost port=5432 user=user password=password dbname=mydatabase sslmode=disable"
	db2DSN = "host=localhost port=5433 user=user password=password dbname=mydatabase sslmode=disable"
)

func main() {
	db1, err := sql.Open("postgres", db1DSN)
	if err != nil {
		log.Fatalf("Could not connect to db1: %v", err)
	}
	defer db1.Close()

	db2, err := sql.Open("postgres", db2DSN)
	if err != nil {
		log.Fatalf("Could not connect to db2: %v", err)
	}
	defer db2.Close()

	userID := 5

	var db *sql.DB
	if userID%2 == 0 {
		db = db1
		fmt.Println("Using database 1")
	} else {
		db = db2
		fmt.Println("Using database 2")
	}

	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS example (id serial PRIMARY KEY, name VARCHAR(50))"); err != nil {
		log.Fatalf("Could not create table: %v", err)
	}
	fmt.Println("Table created or already exists")
}
