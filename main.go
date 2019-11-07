package main

import (
	"html/template"
	"net/http"
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
