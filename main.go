package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os/exec"
)

func main() {
	// Sets up a file server in current directory
	http.HandleFunc("/", index)
	http.HandleFunc("/doctorlogin.html", doctorloginfunc)
	http.HandleFunc("/doctor.html", doctorfunc)
	http.HandleFunc("/patientlogin.html", patientloginfunc)
	http.HandleFunc("/patient.html", patientfunc)
	http.HandleFunc("/employeelogin.html", employeeloginfunc)
	http.HandleFunc("/employee.html", employeefunc)
	http.HandleFunc("/stock.html", stockfunc)
	http.HandleFunc("/prescription.html", prescriptionfunc)
	http.ListenAndServe(":9000", nil)
}

//index runs the index page
func index(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("html/index.html")
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	temp.Execute(response, nil)
}

func doctorloginfunc(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("html/doctorlogin.html")
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	temp.Execute(response, nil)
	//values of form text boxes
	fname := request.FormValue("fname")
	lname := request.FormValue("lname")
	dpass := request.FormValue("dpass")
	//dsubmit := request.FormValue("dsubmit")		//don't know how to get it to run only when submit is pressed
	if fname == "Bob" && lname == "Builder" && dpass == "yes" { //change to compare to database
		fmt.Println("Welcome", fname, lname)   //just for testing
		fmt.Println("Your password is", dpass) //just for testing
		//fmt.Println(dsubmit)

		//add flags to make it check the Doctors table for user info in this case --doc........will need to change this code for other machines
		syst, err := exec.Command("ssh", "user1@192.168.56.102", "cd", "test", ";", "go", "run", "main.go", "--doc").Output()
		if err != nil {
			fmt.Println(err)
		}
		result := string(syst)
		fmt.Println(result)
	} else {
		fmt.Println("wrong") //just for testing
		//fmt.Println(dsubmit)

		//trying to print out when invalid info is entered
		//t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
		//err = t.ExecuteTemplate(out, "T", "<script>alert('your information was incorrect please try again')</script>")
	}
}

func doctorfunc(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("html/doctor.html")
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	temp.Execute(response, nil)
}

func patientloginfunc(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("html/patientlogin.html")
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	temp.Execute(response, nil)
}

func patientfunc(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("html/patient.html")
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	temp.Execute(response, nil)
}

func employeeloginfunc(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("html/employeelogin.html")
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	temp.Execute(response, nil)
}

func employeefunc(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("html/employee.html")
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	temp.Execute(response, nil)
}

func stockfunc(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("html/stock.html")
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	temp.Execute(response, nil)
}

func prescriptionfunc(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("html/prescription.html")
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	temp.Execute(response, nil)
}

//Other machines code --- problem how to send variables and values over ssh?

// package main
// import (
// 	"fmt"
// 	"flag"
// )
// func main {
// docdb := flag.Bool("doc", false, "Check doctor table")
// flag.Parse()
// if docdb == true
// fmt.Println access doc table data //put code to check db
// }
