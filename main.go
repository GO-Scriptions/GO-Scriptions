package main

import (
		"fmt"
		"database/sql"
		"math/rand"

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
	printDoctorTable(db)
	printPharmacists(db)
}

func ping(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

//this function prints the pharmacists Table
func printPharmacists(db *sql.DB) {
	rows, _ := db.Query(`SELECT * FROM Pharmacists`)
	for rows.Next() {
		var firstname string
		var lastname string
		var employeeId string
		var password string
		var isManager string

		rows.Scan(&firstname, &lastname, &employeeId, &password, &isManager)
		fmt.Println(firstname, lastname, employeeId, password, isManager)
	}
}

//This function Prints the Doctors Table
func printDoctorTable(db *sql.DB) {
        rows, _ := db.Query(`SELECT * FROM Doctors`)
        for rows.Next() {
                var First_Name string
                var Last_Name string
                var Phone string
                var Email string
                var Username string
                var Doc_Password string

                rows.Scan(&First_Name, &Last_Name, &Phone, &Email, &Username, &Doc_Password)
		fmt.Println(First_Name, Last_Name, Phone, Email, Username, Doc_Password)
        }
}

// this function creates a new bank account
// Since the pharmacist is the one entering doctors in there, they should not be 
// setting their password for them
func createDoctor(db *sql.DB) {
	var doctorFirst string
	fmt.Println("Doctor First Name")
	fmt.Scan(&doctorFirst)

	var doctorLast string
	fmt.Println("Doctor Last Name:")
	fmt.Scan(&doctorLast)

	var phone string
	fmt.Println("Doctor Last Phone Number:")
	fmt.Scan(&phone)

	var email string
	fmt.Println("Doctor Email:")
	fmt.Scan(&email)

	doctorID := getdoctorID()

	sqlStatement := `INSERT INTO Doctors values($1,$2,$3,$4,$5)`
	_, err := db.Exec(sqlStatement, doctorFirst, doctorLast, phone, email, doctorID)

	if err != nil {
		fmt.Print("An error occurred creating this account.", err)
	}
}

// generates a pseudo-random number
func getdoctorID() string {
	rand.Seed(time.Now().UnixNano())
	var num int
	num = rand.Intn(10000)
	fmt.Println(num)

	acctnumstr := strconv.Itoa(num)
	return acctnumstr
}

