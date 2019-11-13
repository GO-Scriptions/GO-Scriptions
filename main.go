package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func main(){
	var name, pw string
	//dockerstart := exec.Command("docker", "start", "postgres")
	//dockerstart.Run()
	name = os.Args[1]
	pw = os.Args[2]
	checkLoginInfo(name, pw)
}

// connect to database
func connect() *sql.DB {
	var conn string
	conn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}
	//fmt.Println("successfully connected to database")
	return db
}
// verify login information
func checkLoginInfo(name string, pw string) {
	db := connect()
	// check if in doctors or pharmacists table
	if isDoctor(db,name){
		// check if password matches name
		if passwordMatchesDoctor(db,name,pw){
			fmt.Println("isDoctor")	
		} else {
			fmt.Println("invalidLogin")
		}	
	}else if isPharmacist(db,name){
		// check if password matches name
                if passwordMatchesPharmacist(db,name,pw){
                        fmt.Println("isPharmacist")
                } else {
                        fmt.Println("invalidLogin")
                }
	}else{
		fmt.Println("invalidLogin")
	}
	defer db.Close()
}
// check if username exists in doctor's table
func isDoctor(db *sql.DB, name string) bool{
	rows, _ := db.Query("SELECT username FROM doctors")
	for rows.Next() {
		var username string
		rows.Scan(&username)
		if name == username {
			return true
		}
	}
	return false
}
// check if username exists in pharmacist's table
func isPharmacist(db *sql.DB, name string) bool{
        rows, _ := db.Query("SELECT username FROM pharmacists")
        for rows.Next() {
                var username string
                rows.Scan(&username)
                if name == username {
                        return true
                }
        }
        return false
}
// check if password of doctor matches username
func passwordMatchesDoctor(db *sql.DB, name string, password string) bool {
	var pw string
	row := db.QueryRow("SELECT password FROM doctors WHERE username = $1", name)
	row.Scan(&pw)
	if password == pw {
		return true
	}
	return false
}
// check if password of pharmacist matches username
func passwordMatchesPharmacist(db *sql.DB, name string, password string) bool {
        var pw string
        row := db.QueryRow("SELECT password FROM pharmacists WHERE username = $1", name)
        row.Scan(&pw)
        if password == pw {
                return true
        }
        return false
}




