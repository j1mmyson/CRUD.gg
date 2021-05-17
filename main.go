package main

import (
	"database/sql"
	"embed"
	"fmt"
	"net/http"
	"text/template"
	"time"

	"github.com/google/uuid"

	_ "github.com/go-sql-driver/mysql"
)

// var db *sql.DB
// var tpl *template.Template
var (
	db               *sql.DB
	tpl              *template.Template
	dbSessionCleaned time.Time
)

//go:embed web
var content embed.FS

const sessionLength int = 60

func init() {
	tpl = template.Must(template.ParseGlob("web/templates/*.gohtml"))
	dbSessionCleaned = time.Now()
}

func main() {
	fmt.Println("Head")
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", user, password, host, database)
	var err error
	fmt.Println("connection check..")
	// Connect to mysql server
	db, err = sql.Open("mysql", connectionString)
	fmt.Println("Connecting to DB..")
	checkError(err)
	defer db.Close()
	pingDB(db)
	fmt.Println("Connected to DB")

	http.HandleFunc("/", login)
	http.HandleFunc("/signup", signUp)
	http.HandleFunc("/index", index)
	http.HandleFunc("/logout", logout)
	http.Handle("/web/", http.FileServer(http.FS(content)))
	fmt.Println("Listening...ss")
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	u := getUser(w, req)
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func login(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/index", http.StatusSeeOther)
		return
	}

	if req.Method == http.MethodPost {
		user, err := ReadUser(db, req)
		if err != nil {
			errMsg := map[string]interface{}{"error": err}
			tpl.ExecuteTemplate(w, "login.gohtml", errMsg)
			return
		}
		sID := uuid.New()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
		CreateSession(db, c.Value, user.Id)
		http.Redirect(w, req, "/index", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "login.gohtml", nil)
}

func signUp(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/index", http.StatusSeeOther)
		return
	}
	if req.Method == http.MethodGet {
		tpl.ExecuteTemplate(w, "signup.gohtml", nil)
	}

	if req.Method == http.MethodPost {
		err := CreateUser(db, req)
		if err != nil {
			errMsg := map[string]interface{}{"error": err}
			tpl.ExecuteTemplate(w, "signup.gohtml", errMsg)
		} else {
			http.Redirect(w, req, "/", http.StatusSeeOther)
		}
		return
	}
}

func logout(w http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	c, _ := req.Cookie("session")
	// delete session
	DeleteSession(db, c.Value)

	//
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)

	if time.Now().Sub(dbSessionCleaned) > (time.Second * 30) {
		go CleanSessions(db)
	}

	http.Redirect(w, req, "/", http.StatusSeeOther)
}
