package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func getUser(w http.ResponseWriter, req *http.Request) User {
	fmt.Println("getUser()")
	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		sID := uuid.New()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
	}
	c.MaxAge = sessionLength
	http.SetCookie(w, c)

	// if the user exists already, get user
	var u User

	un, err := ReadSession(db, c.Value)
	if err != nil {
		log.Fatal(err)
	}
	UpdateCurrentTime(db, un)
	u, _ = ReadUserById(db, un)
	return u
}

func alreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	fmt.Println("alreadyLoggedIn()")
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}

	un, err := ReadSession(db, c.Value)
	if err != nil {
		return false
	}

	UpdateCurrentTime(db, un)

	_, err = ReadUserById(db, un)
	if err != nil {
		return false
	}

	c.MaxAge = sessionLength
	http.SetCookie(w, c)
	return true
}
