package cmd

import (
	"log"
	"net/http"
)

type Role int

var (
	Admin     Role = 0
	Moderator Role = 1
	User      Role = 2
	Guest     Role = 3
)

func Authorize(next http.Handler, role Role) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(">> role:", role, r.URL.Path)
		if Admin != role {
			// todo error
		}
		next.ServeHTTP(w, r)

	})
}
