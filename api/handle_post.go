package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	model "../model"
)

func AllPostsHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)
	if req.URL.Path != "/api/v1/posts" {
		http.Error(w, "Go back to the main page", 404)
		return
	}
	switch req.Method {
	case "GET":
		ApiPosts_GET(w, req)
	default:
		http.Error(w, "Only GET method allowed, return to main page", 405)
	}
}

func LikePostHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)
	if req.URL.Path != "/api/v1/like-post" {
		http.Error(w, "Go back to the main page", 404)
		return
	}
	switch req.Method {
	case "GET":
		LikePost_GET(w, req)
	default:
		http.Error(w, "Only GET method allowed, return to main page", 405)
	}
}

func DislikePostHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)
	if req.URL.Path != "/api/v1/dislike-post" {
		http.Error(w, "Go back to the main page", 404)
		return
	}
	switch req.Method {
	case "GET":
		DislikePost_GET(w, req)
	default:
		http.Error(w, "Only GET method allowed, return to main page", 405)
	}
}

func ApiPosts_GET(w http.ResponseWriter, r *http.Request) {
	limit := r.URL.Query()["limit"]
	offset := r.URL.Query()["offset"]

	if limit == nil || len(limit) != 1 || offset == nil || len(offset) != 1 {
		// todo render 404
		http.Error(w, "Params required", 404)
		return
	}

	limitInt, err := strconv.Atoi(limit[0])
	if err != nil {
		http.Error(w, "Invalid limit param", 404)
		return
	}
	offsetInt, err := strconv.Atoi(offset[0])
	if err != nil {
		http.Error(w, "Id required", 404)
		return
	}

	posts, err := model.GlobalPostStore.GetLatest(offsetInt, limitInt)
	if err != nil || posts == nil {
		fmt.Println(err)
		http.Error(w, "Error", 404)
		return
	}

	for _, post := range posts {
		if len(post.Text) < 400 {
			continue
		}
		post.Text = post.Text[:400]
	}

	js, err2 := json.Marshal(posts)
	if err2 != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func LikePost_GET(w http.ResponseWriter, r *http.Request) {
	postIDArr := r.URL.Query()["postId"]
	if postIDArr == nil || len(postIDArr) != 1 {
		// todo render 404
		http.Error(w, "Params required", 404)
		return
	}

	session := model.RequestSession(r)
	if session == nil {
		w.WriteHeader(http.StatusUnauthorized)
	}

	postID, err := strconv.Atoi(postIDArr[0])
	if err != nil {
		http.Error(w, "Invalid limit param", 404)
		return
	}

	reactHelper(postID, session, model.LIKE_COUNT, model.LIKE)
	post, err := model.GlobalPostStore.Get(int64(postID))
	js, err2 := json.Marshal(post)
	if err2 != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func DislikePost_GET(w http.ResponseWriter, r *http.Request) {
	postIDArr := r.URL.Query()["postId"]
	if postIDArr == nil || len(postIDArr) != 1 {
		// todo render 404
		http.Error(w, "Params required", 404)
		return
	}

	session := model.RequestSession(r)
	if session == nil {
		w.WriteHeader(http.StatusUnauthorized)
	}

	postID, err := strconv.Atoi(postIDArr[0])
	if err != nil {
		http.Error(w, "Invalid limit param", 404)
		return
	}

	reactHelper(postID, session, model.DISLIKE_COUNT, model.DISLIKE)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	post, err := model.GlobalPostStore.Get(int64(postID))
	js, err2 := json.Marshal(post)
	if err2 != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func reactHelper(postID int, session *model.Session, reactName string, reactCode int) error {
	otherReactCode, otherReactName := model.DISLIKE, model.DISLIKE_COUNT
	if reactName == model.DISLIKE_COUNT {
		otherReactCode, otherReactName = model.LIKE, model.LIKE_COUNT
	}

	hasReacted, err := model.GlobalPostStore.HasReacted(postID, session.UserID, reactCode)
	if err != nil {
		return err
	}

	if !hasReacted {
		err = model.GlobalPostStore.IncrementReaction(int64(postID), reactName)
		if err != nil {
			return err
		}

		err = model.GlobalPostStore.NewUserReaction(postID, session.UserID, reactCode)
		if err != nil {
			return err
		}
	} else {
		err = model.GlobalPostStore.DecrementReaction(int64(postID), reactName)
		if err != nil {
			return err
		}
	}

	hasReacted, err = model.GlobalPostStore.HasReacted(postID, session.UserID, otherReactCode)
	if err != nil {
		return err
	}
	if hasReacted {
		err = model.GlobalPostStore.DecrementReaction(int64(postID), otherReactName)
	}

	return nil

}
