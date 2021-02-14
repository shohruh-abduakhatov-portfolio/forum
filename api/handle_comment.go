package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	model "forum.com/model"
)

func PostCommentsHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)
	if req.URL.Path != "/api/v1/comments" {
		http.Error(w, "Go back to the main page", 404)
		return
	}
	switch req.Method {
	case "GET":
		PostComments_GET(w, req)
	default:
		http.Error(w, "Only GET method allowed, return to main page", 405)
	}
}

func NewCommentHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)
	if req.URL.Path != "/api/v1/comment/new" {
		http.Error(w, "Go back to the main page", 404)
		return
	}
	switch req.Method {
	case "POST":
		NewComment_POST(w, req)
	default:
		http.Error(w, "Only GET method allowed, return to main page", 405)
	}
}

func PostComments_GET(w http.ResponseWriter, r *http.Request) {
	postID := r.URL.Query()["postId"]
	if postID == nil || len(postID) != 1 {
		// todo render 404
		http.Error(w, "Params required", 404)
		return
	}

	id, err := strconv.Atoi(postID[0])
	if err != nil {
		http.Error(w, "Invalid limit param", 404)
		return
	}

	comments, err := model.GlobalCommentStore.GetByTopic(int64(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, err2 := json.Marshal(comments)
	if err2 != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func NewComment_POST(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/comment/new")
	comment := r.FormValue("comment")
	postIDStr := r.FormValue("postId")

	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		http.Error(w, "Invalid limit param", http.StatusBadRequest)
		return
	}

	user := model.RequestUser(r)
	newComment, err := model.NewComment(user, int64(postID), comment)
	if err != nil {
		http.Error(w, "could not create new comment", http.StatusInternalServerError)
		return
	}

	// send response
	js, err2 := json.Marshal(newComment)
	if err2 != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
