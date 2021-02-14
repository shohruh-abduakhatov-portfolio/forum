package internal

import (
	"log"
	"net/http"
	"os"

	model "forum.com/model"
	render "forum.com/render"
)

func AuthenticateRequest(w http.ResponseWriter, r *http.Request) {
	// Redirect the user to login if they’re not authenticated
	authenticated := false
	if !authenticated {
		http.Redirect(w, r, "/auth", http.StatusFound)
	}
}

func LoginHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/login" {
		http.Error(w, "Go back to the main page", 404)
		return
	}
	session := model.RequestSession(req)
	if session != nil {
		http.Redirect(w, req, "/posts", http.StatusFound)
	}
	switch req.Method {
	case "GET":
		HandleLogin_GET(w, req)
	case "POST":
		HandleLogin_POST(w, req)
	default:
		http.Error(w, "Only GET method allowed, return to main page", 405)
	}
}

func HandleRegister(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/register" {
		http.Error(w, "Go back to the main page", 404)
		return
	}
	session := model.RequestSession(req)
	if session != nil {
		http.Redirect(w, req, "/posts", http.StatusFound)
	}
	switch req.Method {
	case "GET":
		HandleRegister_GET(w, req)
	case "POST":
		HandleRegister_POST(w, req)
	default:
		http.Error(w, "Only GET method allowed, return to main page", 405)
	}
}

func HandleLogin_GET(w http.ResponseWriter, r *http.Request) {
	next := r.URL.Query().Get("next")
	render.Basic(w, r, "login.html", map[string]interface{}{
		"Next":     next,
		"ClientID": os.Getenv("GithubClientID"),
		"Host":     os.Getenv("Host"),
		"Port":     os.Getenv("Port"),
		"Protocol": os.Getenv("Protocol"),
	})

	// w.Header().Set("Content-Type", "text/html")
	// render.Basic(w, req, "login.html", nil)
}

func HandleLogin_POST(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	next := r.FormValue("next")
	rememberMe := r.FormValue("rememberMe")

	user, err := model.FindUser(username, password)
	if err != nil {
		if model.IsValidationError(err) {
			render.Basic(w, r, "login.html", map[string]interface{}{
				"Error": err,
				"User":  user,
				"Next":  next,
			})
			return
		}
		panic(err)
	}
	w.Header().Set("timeout", rememberMe)
	session := model.FindOrCreateSession(w, r)
	session.UserID = user.ID
	err = model.GlobalSessionStore.Save(session)
	if err != nil {
		if !model.IsSqliteError(err) && !model.IsUniqueConstraintError(err) {
			panic(err)
		}

	}
	if next == "" {
		next = "/posts"
	}
	cookie, err := r.Cookie("GophrSession")
	if err == nil {
		log.Println("/", r.URL.Path)
		log.Println(">>> ", cookie.Value)
	}

	http.Redirect(w, r, next+"?flash=Signed+in", http.StatusFound)
}

func HandleRegister_GET(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	render.Basic(w, req, "register.html", map[string]interface{}{
		"ClientID": os.Getenv("GithubClientID"),
		"Host":     os.Getenv("Host"),
		"Port":     os.Getenv("Port"),
		"Protocol": os.Getenv("Protocol"),
	})
}

func HandleRegister_POST(w http.ResponseWriter, r *http.Request) {
	user, err := model.NewUser(
		r.FormValue("username"),
		r.FormValue("email"),
		r.FormValue("password"),
	)
	retype := r.FormValue("retypePassword")
	pass := r.FormValue("password")
	if retype != pass {
		render.Basic(w, r, "register.html", map[string]interface{}{
			"Error": "Passwords did not match",
			"User":  user,
		})
		return
	}
	if err != nil {
		if model.IsValidationError(err) {
			render.Basic(w, r, "register.html", map[string]interface{}{
				"Error": err.Error(),
				"User":  user,
			})
			return
		}
		http.Error(w, "Go back to the main page", 404)
		return
	}
	// save
	_, err = model.GlobalUserStore.New(user)
	if err != nil {
		render.Basic(w, r, "register.html", map[string]interface{}{
			"Error": err.Error(),
			"User":  user,
		})
		return
	}
	// Create a new session
	session := model.NewSession(w)
	session.UserID = user.ID
	err = model.GlobalSessionStore.Save(session)
	if err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/posts?flash=User+created", http.StatusFound)
}
