package cmd

import (
	"log"
	"net/http"
	"net/url"

	model "forum.com/model"
)

type Role int

// 777 - READ, WRITE, TOTAL_CONTROL
var (
	Admin     Role = 777
	Moderator Role = 770
	User      Role = 0
	Guest     Role = -777
)

func Authorize(next http.Handler, role Role) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(">> role:", role, r.URL.Path)
		user := model.RequestUser(r)
		if user == nil || user.RoleID < int(role) {
			log.Println(">> Has access: false")
			// Save current path to serve after login
			query := url.Values{}
			query.Add("next", url.QueryEscape(r.URL.String()))
			// Redirect to login
			http.Redirect(w, r, "/login?"+query.Encode(), http.StatusFound)
			return
		}
		log.Println(">> Has access: true")
		next.ServeHTTP(w, r)

	})
}
