package internal

import (
	"net/http"

	model "forum.com/model"
)

func HandleSignOut(w http.ResponseWriter, r *http.Request) {
	session := model.RequestSession(r)
	if session != nil {
		err := model.GlobalSessionStore.Delete(session)
		if err != nil {
			panic(err)
		}
	}
	// render.Basic(w, r, "login.html", nil)
	http.Redirect(w, r, "/login?flash=Signed+out", http.StatusFound)
}
