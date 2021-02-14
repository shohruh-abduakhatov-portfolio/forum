package internal

import (
	"fmt"
	"net/http"
)

func Handle(w http.ResponseWriter, req *http.Request) {
	fmt.Println("/")
	if req.URL.Path != "/" {
		http.Error(w, "Go back to the main page", 404)
		return
	}

	switch req.Method {
	case "GET":
		http.Redirect(w, req, "/posts", http.StatusFound)
	default:
		http.Error(w, "Only GET method allowed, return to main page", 405)
	}
}
