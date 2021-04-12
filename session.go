package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func getUser(w http.ResponseWriter, req *http.Request) User {
	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		sID := uuid.New()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}

	}
	http.SetCookie(w, c)

	// if the user exists already, get user
	var u User

	un, err := ReadSession(db, c.Value)
	fmt.Println("2 un = ", un)
	fmt.Println(err)
	if err != nil {
		log.Fatal(err)
	}
	u, _ = ReadUserById(db, un)
	fmt.Println("user in getUser() = ", u)
	return u
}

func alreadyLoggedIn(req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}

	un, _ := ReadSession(db, c.Value)
	_, err = ReadUserById(db, un)
	// un := dbSessions[c.Value]
	// _, ok := dbUsers[un]
	if err != nil {
		return false
	}
	return true
}
