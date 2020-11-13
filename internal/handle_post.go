package internal

import (
	"fmt"
	"net/http"
	"strconv"

	model "../model"
	render "../render"
	utils "../utils"
)

func AllPostsHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/posts" {
		http.Error(w, "Go back to the main page", 404)
		return
	}
	switch req.Method {
	case "GET":
		HandlePosts_GET(w, req)
	default:
		http.Error(w, "Only GET method allowed, return to main page", 405)
	}
}

func PostHandler(w http.ResponseWriter, req *http.Request) {
	id, ok := getCode(req, "/post/")
	if !ok {
		http.Error(w, "Go back to the main page", 404)
		return
	}
	switch req.Method {
	case "GET":
		HandlePost_GET(w, req, id)
	default:
		http.Error(w, "Only GET method allowed, return to main page", 405)
	}
}

func AllPostsCategoryHandler(w http.ResponseWriter, req *http.Request) {
	id, ok := getCode(req, "/posts-category/")
	fmt.Println(id, ok)
	if !ok {
		http.Error(w, "Go back to the main page", 404)
		return
	}
	switch req.Method {
	case "GET":
		HandlePostsCategory_GET(w, req, id)
	default:
		http.Error(w, "Only GET method allowed, return to main page", 405)
	}
}

func NewPostHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/post/new" {
		http.Error(w, "Go back to the main page", 404)
		return
	}
	switch req.Method {
	case "GET":
		HandleNewPost_GET(w, req)
	case "POST":
		HandleNewPost_POST(w, req)
	default:
		http.Error(w, "Only GET method allowed, return to main page", 405)
	}
}

func NewEditHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/post/edit" {
		http.Error(w, "Go back to the main page", 404)
		return
	}
	switch req.Method {
	case "GET":
		HandleEditPost_GET(w, req)
	case "POST":
		HandleEditPost_POST(w, req)
	default:
		http.Error(w, "Only GET method allowed, return to main page", 405)
	}
}

func NewDeleteHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/post/delete" {
		http.Error(w, "Go back to the main page", 404)
		return
	}
	switch req.Method {
	case "DELETE":
		HandleDeletePost_DELETE(w, req)
	default:
		http.Error(w, "Only GET method allowed, return to main page", 405)
	}
}

func HandlePosts_GET(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/posts - GET")
	render.Template(w, r, "post/posts", map[string]interface{}{
		"Mode": "All Posts.",
	})
}

func HandlePost_GET(w http.ResponseWriter, r *http.Request, id int64) {
	fmt.Println("/post - GET")

	session := model.RequestSession(r)

	post, err := model.GlobalPostStore.Get(id)
	if err != nil || post == nil {
		http.Error(w, "Error", 404)
		return
	}

	render.Template(w, r, "post/view", map[string]interface{}{
		"hasAccess": session != nil,
		"Post":      post,
		"likes":     abbr(post.LikeCount),
		"dislikes":  abbr(post.DislikeCount),
		"comments":  abbr(post.CommentCount),
		"IsMyPost":  post.UserID == session.UserID,
	})
}

func HandlePostsCategory_GET(w http.ResponseWriter, r *http.Request, id int64) {
	fmt.Println("/posts - GET")
	render.Template(w, r, "post/by-category", map[string]interface{}{
		"Mode":       "Posts By Category",
		"CategoryId": id,
	})
}

func HandleNewPost_GET(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "post/new", map[string]interface{}{
		"Action": "New",
	})
}

func HandleEditPost_GET(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query()["id"]
	if id == nil || len(id) != 1 {
		// todo render 404
		http.Error(w, "Id required", 404)
		return
	}

	postId, err := strconv.Atoi(id[0])
	if err != nil {
		http.Error(w, "Id required", 404)
		return
	}

	post, err := model.GlobalPostStore.Get(int64(postId))
	if err != nil || post == nil {
		http.Error(w, "Id required", 404)
		return
	}

	render.Template(w, r, "post/new", map[string]interface{}{
		"Action": "Edit",
		"Post":   post,
	})
}

func HandleNewPost_POST(w http.ResponseWriter, r *http.Request) {
	imgPath, err := utils.ParseImage(w, r)
	if err != nil {
		if !model.IsValidationError(err) {
			err = model.ErrWentWrong()
		}
		render.Template(w, r, "post/new", map[string]interface{}{
			"Error": err.Error(),
		})
		return
	}

	currentUser := model.RequestUser(r)
	title := r.FormValue("title")
	text := r.FormValue("text")

	// Validate
	post, err := model.NewPost(currentUser, title, text, imgPath)
	if err != nil {
		if !model.IsValidationError(err) {
			err = model.ErrWentWrong()
		}
		render.Template(w, r, "post/new", map[string]interface{}{
			"Error": err.Error(),
		})
		return
	}

	// Save post
	_, err = model.GlobalPostStore.New(post)
	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, "/account?flash=Post+Created", http.StatusFound)
}

func HandleEditPost_POST(w http.ResponseWriter, r *http.Request) {
	imgPath, err := utils.ParseImage(w, r)
	if err != nil {
		if !model.IsValidationError(err) {
			err = model.ErrWentWrong()
		}
		render.Template(w, r, "post/new", map[string]interface{}{
			"Error": err.Error(),
		})
		return
	}

	currentUser := model.RequestUser(r)
	title := r.FormValue("title")
	text := r.FormValue("text")

	// Validate
	post, err := model.NewPost(currentUser, title, text, imgPath)
	if err != nil {
		if !model.IsValidationError(err) {
			err = model.ErrWentWrong()
		}
		render.Template(w, r, "post/new", map[string]interface{}{
			"Error": err.Error(),
		})
		return
	}

	// Save post
	_, err = model.GlobalPostStore.New(post)
	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, "/account?flash=Post+Created", http.StatusFound)
}

func HandleDeletePost_DELETE(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query()["id"]
	if id == nil || len(id) != 1 {
		// todo render 404
		http.Error(w, "Id required", 404)
		return
	}
	postId, err := strconv.Atoi(id[0])
	if err != nil {
		http.Error(w, "Id required", 404)
		return
	}
	err = model.GlobalPostStore.Delete(int64(postId))
	if err != nil {
		http.Error(w, "Id required", 404)
		return
	}

	http.Redirect(w, r, "/account?flash=Post+Deleted", http.StatusFound)
}

func abbr(count int64) string {
	if count > 999999 {
		return fmt.Sprintf("%v M", count/1000000)
	} else if count > 999 {
		return fmt.Sprintf("%v K", count/1000)
	} else {
		return fmt.Sprintf("%v", count)
	}
}
