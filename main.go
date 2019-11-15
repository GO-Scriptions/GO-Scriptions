package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "127.0.0.1"
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
	seePrescriptions(db)
}

func ping(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

//This function Prints the Doctors Table
func printDoctorTable(db *sql.DB) {
	rows, _ := db.Query(`SELECT * FROM Doctors`)
	for rows.Next() {
		var firstName string
		var lastName string
		var docUsername string
		var docPassword string

		rows.Scan(&firstName, &lastName, &docUsername, &docPassword)
		fmt.Println(firstName, lastName, docUsername, docPassword)
	}
}

//this function prints the pharmacists Table
func printPharmacists(db *sql.DB) {
	rows, _ := db.Query(`SELECT * FROM Pharmacists`)
	for rows.Next() {
		var firstname string
		var lastname string
		var username string
		var password string
		var isManager string

		rows.Scan(&firstname, &lastname, &username, &password, &isManager)
		fmt.Println(firstname, lastname, username, password, isManager)
	}
}

//this function prints out the inventory table
func printInventory(db *sql.DB) {
	rows, _ := db.Query(`SELECT * FROM Inventory`)
	for rows.Next() {
		var drugname string
		var amtOnHand int
		var costpermg float64
		var supplier string

		rows.Scan(&drugname, &amtOnHand, &costpermg, &supplier)
		fmt.Println(drugname, amtOnHand, costpermg, supplier)
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

	var docUsername string
	fmt.Println("Doctor Username:")
	fmt.Scan(&docUsername)

	var docPassword string
	fmt.Println("Doctor Password")
	fmt.Scan(&docPassword)

	sqlStatement := `INSERT INTO Doctors values($1,$2,$3,$4)`
	_, err := db.Exec(sqlStatement, doctorFirst, doctorLast, docUsername, docPassword)

	if err != nil {
		fmt.Print("An error occurred creating this account.", err)
	}
}

//function to print current  precriptions table
func seePrescriptions(db *sql.DB) {
	rows, _ := db.Query(`SELECT * FROM Prescriptions`)
	for rows.Next() {
		var prescID string
		var docPrescribing string
		var drugName string
		var amount int
		var patientFirst string
		var patientLast string
		var cost float64
		var prescStatus string
		var datePrescribed string

		rows.Scan(&prescID, &docPrescribing, &drugName, &amount, &patientFirst, &patientLast, &cost, &prescStatus, &datePrescribed)
		fmt.Println(prescID, docPrescribing, drugName, amount, patientFirst, patientLast, cost, prescStatus, datePrescribed)
	}
}

//function to create new entries in prescription history table
