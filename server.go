package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func index(res http.ResponseWriter, req *http.Request) {
	tpl := template.Must(template.ParseFiles("template/index.htm"))
	err := tpl.Execute(res, nil)
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}

func signUp(res http.ResponseWriter, req *http.Request) {
	tpl := template.Must(template.ParseFiles("template/signUp.htm"))
	err := tpl.Execute(res, nil)
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/signUp", signUp)
	http.Handle("/template/", http.StripPrefix("/template/", http.FileServer(http.Dir("template"))))
	fmt.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}
