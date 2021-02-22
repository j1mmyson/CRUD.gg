package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	host     = "localhost"
	database = "tutorial"
	user     = "byungwook"
	password = "quddnr!2"
)

// Topic table columns
type Topic struct {
	Id          int
	Title       string
	Description string
	Created     string
	Author      string
	Profile     string
}

// Create1 insert data to db
func Create1(db *sql.DB) {
	// Create 1
	insert, err := db.Query("INSERT INTO topic (title, description, created, author, profile) VALUES ('GOPHER', 'Hello Golang', NOW(), 'byungwook', 'dev')")
	checkError(err)
	defer insert.Close()
}

// Create2 insert data to db
func Create2(db *sql.DB) {
	// Create 2
	stmt, err := db.Prepare("insert into topic (title, description, created, author, profile) values (?, ?, ?, ?, ?)")
	checkError(err)
	defer stmt.Close()
	t := time.Now().Format("2006-01-02 15:04:05")
	res, err := stmt.Exec("golang", "hello golang!", t, "jimmy", "developer")
	checkError(err)
	count, err := res.RowsAffected()
	checkError(err)
	fmt.Println(count)
}

// Read select all data from db
func Read(db *sql.DB) {
	// Read
	rows, err := db.Query("select * from topic")
	checkError(err)
	var topic = Topic{}

	for rows.Next() {
		err = rows.Scan(&topic.Id, &topic.Title, &topic.Description, &topic.Created, &topic.Author, &topic.Profile)
		checkError(err)
		fmt.Println(topic)
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

func main() {
	fmt.Println("Go MYSQL Tutorial")
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", user, password, host, database)

	// Connect to mysql server
	db, err := sql.Open("mysql", connectionString)
	checkError(err)
	defer db.Close()
	pingDB(db)

	Read(db)
	Create2(db)
	Read(db)

}
