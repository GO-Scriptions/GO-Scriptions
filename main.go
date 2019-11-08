package main

import (
    "net/http"
    "html/template"
    "fmt"
    "time"
    "math/rand"

    "github.com/gorilla/sessions"
)
type Info struct{
        Authenticated bool
        Username string
}
var (
    // key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
    //key = []byte("super-secret-key")
    key = makeKey()
    store = sessions.NewCookieStore([]byte(key))
)
func index(response http.ResponseWriter, request *http.Request) {
    temp, _ := template.ParseFiles("index.html")
    session, _ := store.Get(request, "cookie-name")
    session.Values["authenticated"] = false
    session.Save(request, response)
    temp.Execute(response, nil)
}
func test(response http.ResponseWriter, request *http.Request) {
    temp, _ := template.ParseFiles("test.html")
    session, _ := store.Get(request, "cookie-name")
    // Check if user is authenticated
    auth, ok := session.Values["authenticated"]
    if !ok {
        http.Error(response, "Forbidden", http.StatusForbidden)
        return
    }
    info := Info{}
    info.Authenticated = auth.(bool)
    name, ok := session.Values["name"]
    if !ok {
        fmt.Println("no name")
    }else {
        info.Username = name.(string)
    }
    temp.Execute(response, info)
}

func login(response http.ResponseWriter, request *http.Request) {
    temp, _ := template.ParseFiles("login.html")
    session, _ := store.Get(request, "cookie-name")
    name := request.FormValue("name")
    fmt.Println("name:",name)
    // query db to authenticate....

    // Set user as authenticated
    session.Values["authenticated"] = true
    session.Values["name"] = name
    session.Save(request, response)
    temp.Execute(response, nil)
}

func logout(response http.ResponseWriter, request *http.Request) {
    temp, _ := template.ParseFiles("logout.html")
    session, _ := store.Get(request, "cookie-name")

    // Revoke users authentication
    session.Values["authenticated"] = false
    session.Values["name"] = ""
    session.Save(request, response)
    temp.Execute(response, nil)
}

func makeKey() string {
        var i int
        var key string
        var chars = [64]string{"a","b","c","d","e","f","g","h","1","2","3","4","5","6","7","8"}
        for i<16 {
                rand.Seed(time.Now().UnixNano())
                number := rand.Intn(16)
                key += chars[number]
                i++
        }
        fmt.Println("key:",key)
        return key

}

func main() {
    http.HandleFunc("/", index)
    http.HandleFunc("/login", login)
    http.HandleFunc("/logout", logout)
    http.HandleFunc("/test", test)

    http.ListenAndServe(":7000", nil)
}

