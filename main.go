package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func main() {
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", datasource)
	defer db.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println("Hello GO Scriptions!")
	ping(db)
}

func ping(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

//this function prints the doctors Table
func printDoctorTable(db *sql.DB) {
	rows, _ := db.Query(`SELECT * FROM Doctors`)
	for rows.Next() {
		var firstname string
		var lastname string
		var phone string
		var email string
		var username string
		var password string

		rows.Scan(&firstname, &lastname, &phone, &email, &username, &password)
		fmt.Println(firstname, lastname, phone, email, username, password)
	}
}
