package main

import (
	"net/http"

	"github.com/google/uuid"
)

func getUser(w http.ResponseWriter, r *http.Request) user {
	// get cookie
	c, err := r.Cookie("session")
	if err != nil {
		uuid := uuid.New()
		c = &http.Cookie{
			Name:  "session",
			Value: uuid.String(),
		}
	}
	http.SetCookie(w, c)

	// if the user exists already, get user
	var u user
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}
	return u
}

func alreadyLoggedIn(r *http.Request) bool {
	c, err := r.Cookie("session")
	if err != nil {
		return false
	}
	un := dbSessions[c.Value]
	_, ok := dbUsers[un]
	return ok
}
