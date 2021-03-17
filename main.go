package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

func index(res http.ResponseWriter, req *http.Request) {
	tpl := template.Must(template.ParseFiles("template/index.htm"))
	err := tpl.Execute(res, nil)
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}

func signUp(res http.ResponseWriter, req *http.Request) {
	fmt.Println("signUp")
	tpl := template.Must(template.ParseFiles("template/signUp.htm"))
	err := tpl.Execute(res, nil)
	if err != nil {
		log.Fatalln("error executing template", err)
	}

	if req.Method == "POST" {
		fmt.Println(db)
		fmt.Println("Create2 worked")
		Create2(db, req)
	}
}

var db *sql.DB

func main() {
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", user, password, host, database)
	var err error
	// Connect to mysql server
	fmt.Println(db)
	db, err = sql.Open("mysql", connectionString)
	checkError(err)
	defer db.Close()
	pingDB(db)
	fmt.Println(db)
	http.HandleFunc("/", index)
	http.HandleFunc("/signUp", signUp)
	http.Handle("/template/", http.StripPrefix("/template/", http.FileServer(http.Dir("template"))))
	fmt.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}
