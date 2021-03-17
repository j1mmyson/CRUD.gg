package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	host     = "localhost"
	database = "gocrud"
	user     = "byungwook"
	password = "quddnr!2"
)

// Topic table columns
type User struct {
	Id       int
	Password string
	Name     string
	Created  string
}

// Create1 insert data to db
func Create1(db *sql.DB) {
	// Create 1
	insert, err := db.Query("INSERT INTO topic (title, description, created, author, profile) VALUES ('GOPHER', 'Hello Golang', NOW(), 'byungwook', 'dev')")
	checkError(err)
	defer insert.Close()
}

// Create2 insert data to db
func Create2(db *sql.DB, req *http.Request) {
	// req.ParseForm()
	id := req.PostFormValue("id")
	password := req.PostFormValue("password")
	name := req.PostFormValue("name")
	t := time.Now().Format("2006-01-02 15:04:05")
	// Create 2
	stmt, err := db.Prepare("insert into user (id, password, name, created) values (?, ?, ?, ?)")
	checkError(err)
	defer stmt.Close()
	res, err := stmt.Exec(id, password, name, t)
	checkError(err)
	count, err := res.RowsAffected()
	checkError(err)
	fmt.Println(count, "rows affected")
}

// Read select all data from db
func Read(db *sql.DB) {
	// Read
	rows, err := db.Query("select * from topic")
	checkError(err)
	var user = User{}

	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Password, &user.Name, &user.Created)
		checkError(err)
		fmt.Println(user)
	}
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

// Delete delete data from db
func Delete(db *sql.DB) {
	// Delete
	stmt, err := db.Prepare("delete from topic where id=?")
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

	Read(db)
	// Create2(db)
	Read(db)

}
