package main

import (
    "net/http"
    "html/template"
    "strings"
    "fmt"

     // path after ~/go/src/ (may need to modify depending on your directory set up)
    "github.com/project-2/GO-Scriptions/interact"
)
type LoginInfo struct{
	Username string
        Doctor bool
	Pharmacist bool
}
var loginInfo = LoginInfo{} 


func index(response http.ResponseWriter, request *http.Request) {
    temp, _ := template.ParseFiles("templates/index.html")
    temp.Execute(response, loginInfo)
}
func login(response http.ResponseWriter, request *http.Request) {
    temp, _ := template.ParseFiles("templates/login.html")
    var cmd, db_response string
    name := request.FormValue("name")
    pw := request.FormValue("pw")
    // go run main.go
    cmd += "/usr/local/go/bin/go run main.go "
    // add command line arguments
    cmd += name
    cmd += " "
    cmd += pw
    fmt.Println("command:",cmd)
    // get database response
    db_response = interact.ExecuteCommand(cmd)
    db_response = strings.TrimSpace(db_response)
    fmt.Println("db response:",db_response)
    if db_response == "isDoctor" {
	loginInfo.Doctor = true
	loginInfo.Username = name
    } else if db_response == "isPharmacist" {
	loginInfo.Pharmacist = true
	loginInfo.Username = name
    }
    temp.Execute(response, loginInfo)
}
func logout(response http.ResponseWriter, request *http.Request) {
    temp, _ := template.ParseFiles("templates/logout.html")
    // revoke user's authentication
    loginInfo = LoginInfo{}
    temp.Execute(response, nil)
}
func main() {
    // initial connection with db
    interact.Init()     

    http.HandleFunc("/", index)
    http.HandleFunc("/login", login)
    http.HandleFunc("/logout", logout)

    http.ListenAndServe(":80", nil)
}
