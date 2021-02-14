package internal

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	model "forum.com/model"
	render "forum.com/render"
	utils "forum.com/utils"
)

func AllPostsHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/posts" {
		http.Error(w, "Go back to the main page", 404)
		return
	}
	cookie, err := req.Cookie("GophrSession")
	if err == nil {
		log.Println("/", req.URL.Path)
		log.Println(">>> ", cookie.Value)
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

func UserReactedPostHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/user-reacted-posts" {
		http.Error(w, "Go back to the main page", 404)
		return
	}
	switch req.Method {
	case "GET":
		HandleUserReactedPost_GET(w, req)
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
	print(">>> req.URL.Path", req.URL.Path)
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
	model.GlobalSessionStore.GetAllSessions()
	render.Template(w, r, "post/posts", map[string]interface{}{
		"Mode": "All Posts.",
	})
}

func HandleUserReactedPost_GET(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/user-reacted-posts - GET")
	render.Template(w, r, "post/by-liked", map[string]interface{}{
		"Mode": "Liked Posts.",
	})
}

func HandlePost_GET(w http.ResponseWriter, r *http.Request, id int64) {
	fmt.Println("/post - GET")

	session := model.RequestSession(r)
	var userId string = ""
	if session != nil {
		userId = session.UserID
	}

	post, cats, user, err := model.GlobalPostStore.Get(id)
	if err != nil || post == nil {
		fmt.Println(err)
		http.Error(w, "Error", 404)
		return
	}

	render.Template(w, r, "post/view", map[string]interface{}{
		"hasAccess":  session != nil,
		"Post":       post,
		"User":       user,
		"likes":      abbr(post.LikeCount),
		"dislikes":   abbr(post.DislikeCount),
		"comments":   abbr(post.CommentCount),
		"IsMyPost":   user.ID == userId,
		"Categories": cats,
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

	post, cats, _, err := model.GlobalPostStore.Get(int64(postId))
	if err != nil || post == nil {
		http.Error(w, "Id required", 404)
		return
	}

	render.Template(w, r, "post/new", map[string]interface{}{
		"Action":     "Edit",
		"Post":       post,
		"Categories": cats,
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
	r.ParseForm()
	fmt.Printf(">>> form: %+v\n", r.Form)
	categoriesSelected := r.Form["category-list"]
	fmt.Println(categoriesSelected)

	ids, err := model.ParseCategoryArr(categoriesSelected)
	if err != nil {
		render.Template(w, r, "post/new", map[string]interface{}{
			"Error": err.Error(),
		})
		return
	}

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
	postId, err := model.GlobalPostStore.New(post)
	if err != nil {
		panic(err)
	}
	post.ID = postId
	if err = model.GlobalPostStore.NewPostCategory(post, ids); err != nil {
		panic(err)
	}
	if err = model.GlobalPostStore.NewUserPost(post); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/posts?flash=Post+Created", http.StatusFound)
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

	r.ParseForm()

	fmt.Printf(">>> form:  %+v\n", r.Form)
	categoriesSelected := r.Form["category-list"]
	if err = model.ValidCategory(categoriesSelected); err != nil {
		render.Template(w, r, "post/new", map[string]interface{}{
			"Error": err.Error(),
		})
		return
	}
	fmt.Println(categoriesSelected)

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
	ids, err := model.ParseCategoryArr(categoriesSelected)
	if err != nil {
		render.Template(w, r, "post/new", map[string]interface{}{
			"Error": err.Error(),
		})
		return
	}

	// Save post
	if err = model.GlobalPostStore.Modify(post); err != nil {
		panic(err)
	}
	if err = model.GlobalPostStore.DeletePostCategories(categoriesSelected); err != nil {
		panic(err)
	}
	if err = model.GlobalPostStore.NewPostCategory(post, ids); err != nil {
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
