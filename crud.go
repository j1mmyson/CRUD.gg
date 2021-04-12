package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

const (
	host     = "localhost"
	database = "gocrud"
	user     = "byungwook"
	password = "quddnr!2"
)

// CustomError: error type struct
type CustomError struct {
	Code    string
	Message string
}

// Topic table columns
type User struct {
	Id       string
	Password string
	Name     string
	Created  string
}

type Input struct {
	Id       string
	Password string
}

func (e *CustomError) Error() string {
	return e.Code + ", " + e.Message
}

// Create1 insert data to db
func Create1(db *sql.DB) {
	// Create 1
	insert, err := db.Query("INSERT INTO topic (title, description, created, author, profile) VALUES ('GOPHER', 'Hello Golang', NOW(), 'byungwook', 'dev')")
	checkError(err)
	defer insert.Close()
}

func CreateSession(db *sql.DB, sessionId string, userId string) {
	stmt, err := db.Prepare("insert into session values (?, ?)")
	checkError(err)
	defer stmt.Close()
	_, err = stmt.Exec(sessionId, userId)
	checkError(err)

}

// Create2 insert data to db
func CreateUser(db *sql.DB, req *http.Request) {
	// req.ParseForm()
	id := req.PostFormValue("id")
	password := req.PostFormValue("password")
	name := req.PostFormValue("name")
	t := time.Now().Format("2006-01-02 15:04:05")
	// Create 2
	stmt, err := db.Prepare("insert into user (id, password, name, created) values (?, ?, ?, ?)")
	checkError(err)
	defer stmt.Close()

	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	_, err = stmt.Exec(id, bs, name, t)
	checkError(err)
}

func ReadSession(db *sql.DB, sessionId string) (string, error) {
	row, err := db.Query("select user_id from session where session_id = ?", sessionId)
	checkError(err)
	var userId string

	for row.Next() {
		err = row.Scan(&userId)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("userId = ", userId)
	}
	return userId, nil
}

func ReadUserById(db *sql.DB, userId string) (User, error) {
	row, err := db.Query("select * from user where id = ?", userId)
	checkError(err)
	var user = User{}
	for row.Next() {
		err := row.Scan(&user.Id, &user.Password, &user.Created, &user.Name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(user)
	}

	return user, nil
}

// Read select all data from db
func ReadUser(db *sql.DB, req *http.Request) (User, error) {
	// Read
	id, pw := req.PostFormValue("id"), req.PostFormValue("password")
	rows, err := db.Query("select * from user where id = ?", id)
	checkError(err)
	var user = User{}

	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Password, &user.Created, &user.Name)
		checkError(err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pw))
	if err != nil {
		return user, &CustomError{Code: "401", Message: "password uncorrect!"}
	}

	return user, nil
}

// Update change data from db
func Update(db *sql.DB) {
	// Update
	stmt, err := db.Prepare("update topic set profile=? where profile=?")
	checkError(err)

	res, err := stmt.Exec("developer", "dev")
	checkError(err)

	a, err := res.RowsAffected()
	checkError(err)

	fmt.Println(a, "rows in set")
}

func DeleteSession(db *sql.DB, sessionID string) {
	stmt, err := db.Prepare("delete from session where session_id=?")
	checkError(err)

	_, err = stmt.Exec(sessionID)
	checkError(err)

}

// Delete delete data from db
func Delete(db *sql.DB) {
	// Delete
	stmt, err := db.Prepare("delete from user where id=?")
	checkError(err)

	res, err := stmt.Exec(5)
	checkError(err)

	a, err := res.RowsAffected()
	checkError(err)
	fmt.Println(a, "rows in set")
}

func pingDB(db *sql.DB) {
	err := db.Ping()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func crud() {
	fmt.Println("Go MYSQL Tutorial")
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", user, password, host, database)

	// Connect to mysql server
	db, err := sql.Open("mysql", connectionString)
	checkError(err)
	defer db.Close()
	pingDB(db)

}
