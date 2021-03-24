package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

func login(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		fmt.Println("post method in index page;")
		user, err := Read(db, req)
		if err != nil {
			log.Println("Login failed")
		} else {
			tpl := template.Must(template.ParseFiles("template/index.gohtml"))
			err := tpl.Execute(res, user)
			if err != nil {
				log.Println("error executing template", err)
			}
		}
	} else {
		tpl := template.Must(template.ParseFiles("template/login.htm"))
		err := tpl.Execute(res, nil)
		if err != nil {
			log.Fatalln("error executing template", err)
		}
	}
}

func signUp(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		tpl := template.Must(template.ParseFiles("template/signUp.htm"))
		err := tpl.Execute(res, nil)
		if err != nil {
			log.Fatalln("error executing template", err)
		}
	}

	if req.Method == "POST" {
		Create2(db, req)
		http.Redirect(res, req, "/", http.StatusSeeOther)
	}
}

var db *sql.DB

func main() {
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", user, password, host, database)
	var err error
	// Connect to mysql server
	// fmt.Println(db)
	db, err = sql.Open("mysql", connectionString)
	checkError(err)
	defer db.Close()
	pingDB(db)
	// fmt.Println(db)
	http.HandleFunc("/", login)
	http.HandleFunc("/signUp", signUp)
	http.Handle("/template/", http.StripPrefix("/template/", http.FileServer(http.Dir("template"))))
	fmt.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}
