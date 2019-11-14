package main

import (
		"fmt"
		"database/sql"

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
		var username string
		var password string
		var isManager string

		rows.Scan(&firstname, &lastname, &username, &password, &isManager)
		fmt.Println(firstname, lastname, username, password, isManager)
	}
}

//this function prints out the inventory table
func printInventory (db *sql.DB) {
	rows, _ := db.Query(`SELECT * FROM Inventory`)
	for rows.Next() {
		var Drug_Name string
		var Amt_On_Hand int
		var Cost_Per_Mg float64
		var Supplier string

		rows.Scan(&Drug_Name, &Amt_On_Hand, &Cost_Per_Mg, &Supplier)
		fmt.Println(Drug_Name, Amt_On_Hand, Cost_Per_Mg, Supplier)
	}
}

//This function Prints the Doctors Table
func printDoctorTable(db *sql.DB) {
        rows, _ := db.Query(`SELECT * FROM Doctors`)
        for rows.Next() {
            var First_Name string
            var Last_Name string
            var Doc_Username string
            var Doc_Password string

            rows.Scan(&First_Name, &Last_Name, &Doc_Username, &Doc_Password)
			fmt.Println(First_Name, Last_Name, Doc_Username, Doc_Password)
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

	var doc_Username string
	fmt.Println("Doctor Username:")
	fmt.Scan(&doc_Username)

	var Doc_Password string
	fmt.Println("Doctor Password")
	fmt.Scan(&Doc_Password)

	sqlStatement := `INSERT INTO Doctors values($1,$2,$3,$4)`
	_, err := db.Exec(sqlStatement, doctorFirst, doctorLast, doc_Username, Doc_Password)

	if err != nil {
		fmt.Print("An error occurred creating this account.", err)
	}
}

//function to print current  precriptions table
func seePrescriptions(db *sql.DB) {
	rows, _ := db.Query(`SELECT * FROM Prescriptions`)
	for rows.Next() {
		var Presc_ID string
		var Doc_Prescribing string
		var Drug_Name string
		var Amount int
		var Patient_First string
		var Patient_Last string
		var Cost float64
		var Presc_Status string
		var Date_Prescribed string

		rows.Scan(&Presc_ID, &Doc_Prescribing, &Drug_Name, &Amount, &Patient_First, &Patient_Last, &Cost, &Presc_Status, &Date_Prescribed)
		fmt.Println(Presc_ID, Doc_Prescribing, Drug_Name, Amount, Patient_First, Patient_Last, Cost, Presc_Status, Date_Prescribed)
	}
}

//function to create new entries in prescription history table