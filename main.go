package main

import (
    "net/http"
    "html/template"
    "fmt"
    "time"
    "math/rand"

    "github.com/gorilla/sessions"
    "github.com/GO-Scriptions/GO-Scriptions/interact"
)
type Info struct{
        Doctor bool
	Pharmacist bool
        Username string
}
var (
    key = makeKey()
    store = sessions.NewCookieStore([]byte(key))
)
func index(response http.ResponseWriter, request *http.Request) {
    temp, _ := template.ParseFiles("index.html")
    temp.Execute(response, nil)
}
func test(response http.ResponseWriter, request *http.Request) {
    temp, _ := template.ParseFiles("test.html")
    session, _ := store.Get(request, "cookie-name")
    // Check if user is authenticated
    doctor, ok := session.Values["doctor"]
    if !ok {
        response.Write([]byte("You do not have permission to access this page."))
        return
    }
    pharmacist, ok := session.Values["pharmacist"]
    if !ok{
        response.Write([]byte("You do not have permission to access this page."))
        return
    }
    if !doctor.(bool) && !pharmacist.(bool){
        response.Write([]byte("You do not have permission to access this page."))
        return
    }
    info := Info{}
    info.Doctor = doctor.(bool)
    info.Pharmacist = pharmacist.(bool)
    name, ok := session.Values["name"]
    if !ok {
        fmt.Println("no name")
    }else {
        info.Username = name.(string)
    }
    temp.Execute(response, info)
}

func loginDoctor(response http.ResponseWriter, request *http.Request) {
    temp, _ := template.ParseFiles("logindoctor.html")
    name := request.FormValue("name")
    // pw := request.FormValue("pw")
   

    // query db to authenticate ....

    // if authenticated
    session, _ := store.Get(request, "cookie-name")
    session.Values["doctor"] = true
    session.Values["pharmacist"] = false
    session.Values["name"] = name
    session.Save(request, response)

    temp.Execute(response, nil)
}
func loginPharmacist(response http.ResponseWriter, request *http.Request) {
    temp, _ := template.ParseFiles("loginpharmacist.html")
    name := request.FormValue("name")
    // pw := request.FormValue("pw")
    

    // query db to authenticate....

    // if authenticated
    session, _ := store.Get(request, "cookie-name")
    session.Values["doctor"] = false
    session.Values["pharmacist"] = true
    session.Values["name"] = name
    session.Save(request, response)

    temp.Execute(response, nil)
}
func logout(response http.ResponseWriter, request *http.Request) {
    temp, _ := template.ParseFiles("logout.html")
    session, _ := store.Get(request, "cookie-name")

    // Revoke user's authentication
    session.Values["doctor"] = false
    session.Values["pharmacist"] = false
    session.Values["name"] = ""
    session.Save(request, response)
    temp.Execute(response, nil)
}

func makeKey() string {
        var i int
        var key string
        var chars = []string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p"}
	chars = append(chars,"q","r","s","t","u","v","w","x","y","z")
	chars = append(chars,"0","1","2","3","4","5","6","7","8","9","$","&","@","*")
        for i<16 {
                rand.Seed(time.Now().UnixNano())
                number := rand.Intn(40)
                key += chars[number]
                i++
        }
        fmt.Println("key:",key)
        return key

}
func main() {
    http.HandleFunc("/", index)
    http.HandleFunc("/loginpharmacist", loginPharmacist)
    http.HandleFunc("/logindoctor", loginDoctor)
    http.HandleFunc("/logout", logout)
    http.HandleFunc("/test", test)

    http.ListenAndServe(":7000", nil)
}
