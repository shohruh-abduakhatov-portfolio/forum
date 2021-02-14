package internal

import (
	"fmt"
	"net/http"

	model "forum.com/model"
	render "forum.com/render"
)

// func HandleUserEdit(w http.ResponseWriter, r *http.Request) {
// 	user := model.RequestUser(r)
// 	render.Template(w, r, "users/edit", map[string]interface{}{
// 		"User": user,
// 	})
// }

func HandleUserEdit(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/edit" {
		http.Error(w, "Go back to the main page", 404)
		return
	}
	switch req.Method {
	case "GET":
		HandleEdit_GET(w, req)
	case "POST":
		HandleEdit_POST(w, req)
	default:
		http.Error(w, "Only GET method allowed, return to main page", 405)
	}
}

func HandleUserAccount(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/account" {
		http.Error(w, "Go back to the main page", 404)
		return
	}
	switch req.Method {
	case "GET":
		HandleAccount_GET(w, req)
	default:
		http.Error(w, "Only GET method allowed, return to main page", 405)
	}
}

func HandleUserPosts(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/user-posts" {
		http.Error(w, "Go back to the main page", 404)
		return
	}
	switch req.Method {
	case "GET":
		HandleUserPosts_GET(w, req)
	default:
		http.Error(w, "Only GET method allowed, return to main page", 405)
	}
}

func HandleUserPosts_GET(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/user-posts - GET")
	render.Template(w, r, "post/by-user", map[string]interface{}{
		"Mode": "My Posts.",
	})
}

func HandleEdit_GET(w http.ResponseWriter, r *http.Request) {
	user := model.RequestUser(r)
	user = &model.User{
		Username: "username text",
		Email:    "my email",
		Password: "my password",
	}
	render.Template(w, r, "user/edit", map[string]interface{}{
		"User": user,
	})
}

func HandleEdit_POST(w http.ResponseWriter, r *http.Request) {
	currentUser := model.RequestUser(r)
	email := r.FormValue("email")
	currentPassword := r.FormValue("currentPassword")
	newPassword := r.FormValue("newPassword")
	user, err := model.UpdateUser(currentUser, email, currentPassword, newPassword)
	if err != nil {
		if model.IsValidationError(err) {
			render.Template(w, r, "users/edit", map[string]interface{}{
				"Error": err.Error(),
				"User":  user,
			})
			return
		}
		panic(err)
	}
	err = model.GlobalUserStore.Update(*currentUser)
	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, "/account?flash=User+updated", http.StatusFound)
}

func HandleAccount_GET(w http.ResponseWriter, r *http.Request) {
	user := model.RequestUser(r)
	user = &model.User{
		Username: "username text",
		Email:    "my email",
		Password: "my password",
	}
	render.Template(w, r, "user/account", map[string]interface{}{
		"User": user,
	})
}
