package main

import (
    "net/http"
    "html/template"
    "fmt"

    "github.com/GO-Scriptions/GO-Scriptions/interact"
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
    name := request.FormValue("name")
    

    // query db to authenticate....
    temp.Execute(response, nil)
}
func logout(response http.ResponseWriter, request *http.Request) {
    temp, _ := template.ParseFiles("templates/logout.html")

    // Revoke user's authentication

}
func main() {
    http.HandleFunc("/", index)
    http.HandleFunc("/login", login)
    http.HandleFunc("/logout", logout)

    http.ListenAndServe(":80", nil)
}
