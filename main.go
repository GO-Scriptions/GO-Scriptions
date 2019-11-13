package main

import (
    "net/http"
    "html/template"
    "fmt"

     // path after ~/go/src/
    "github.com/sdgrandy/project-2/GO-Scriptions/interact"
)
type LoginInfo struct{
	Username string
        Doctor bool
	Pharmacist bool
}
var loginInfo = LoginInfo{} 


func index(response http.ResponseWriter, request *http.Request) {
    temp, _ := template.ParseFiles("templates/index.html")
    temp.Execute(response, nil)
}
func login(response http.ResponseWriter, request *http.Request) {
    temp, _ := template.ParseFiles("/templates/login.html")
    var cmd string
    name := request.FormValue("name")
    pw := request.FormValue("pw")
    // go run main.go
    cmd += "/usr/local/go/bin/go run main.go "
    // add command line arguments
    cmd += name
    cmd += " "
    cmd += pw
    fmt.Println("command is",cmd)
    interact.ExecuteCommand(cmd)
    // query db to authenticate....
    temp.Execute(response, nil)
}
func logout(response http.ResponseWriter, request *http.Request) {
    temp, _ := template.ParseFiles("templates/logout.html")
    // revoke user's authentication
    loginInfo = LoginInfo{}
    temp.Execute(response, nil)
}
func main() {
    http.HandleFunc("/", index)
    http.HandleFunc("/login", login)
    http.HandleFunc("/logout", logout)

    http.ListenAndServe(":8000", nil)
}
