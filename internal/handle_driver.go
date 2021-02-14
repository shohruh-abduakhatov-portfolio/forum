package internal

import (
	"net/http"

	render "forum.com/render"
)

func DriverLogHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/log" {
		http.Error(w, "Go back to the main page", 404)
		return
	}
	switch req.Method {
	case "GET":
		HandleDriverLog_GET(w, req)
	default:
		http.Error(w, "Only GET method allowed, return to main page", 405)
	}
}

func HandleDriverLog_GET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	render.Basic(w, r, "daily_log.html", map[string]interface{}{
		"title": "hello",
		"text":  "world",
	})
}
