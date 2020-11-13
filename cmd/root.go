package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"

	api "../api"
	config "../config"
	internal "../internal"
	model "../model"
)

func init() {
	config.InitConfig()
	model.Init()
}

const api_str = "/api/v1"

func Execute() {
	// Api - V1
	http.Handle("/log", http.HandlerFunc(internal.DriverLogHandler))
	// posts
	http.Handle(api_str+"/posts", http.HandlerFunc(api.AllPostsHandler))
	http.Handle(api_str+"/posts-by-category", http.HandlerFunc(api.AllPostsByCategoryHandler))
	http.Handle(api_str+"/like-post", http.HandlerFunc(api.LikePostHandler))
	http.Handle(api_str+"/dislike-post", http.HandlerFunc(api.DislikePostHandler))
	// comments
	http.Handle(api_str+"/comments", http.HandlerFunc(api.PostCommentsHandler))
	http.Handle(api_str+"/comment/new", http.HandlerFunc(api.NewCommentHandler))
	// Categories
	http.Handle(api_str+"/categories", http.HandlerFunc(api.CategoryListHandler))

	// Posts
	http.Handle("/post/new", http.HandlerFunc(internal.NewPostHandler))
	http.Handle("/post/edit", http.HandlerFunc(internal.NewEditHandler))
	http.Handle("/post/delete", http.HandlerFunc(internal.NewEditHandler))
	http.Handle("/posts", http.HandlerFunc(internal.AllPostsHandler))
	http.Handle("/posts-category/", http.HandlerFunc(internal.AllPostsCategoryHandler))
	http.Handle("/post/", http.HandlerFunc(internal.PostHandler))
	http.Handle("/user-reacted-posts", Authorize(http.HandlerFunc(internal.HandleUserAccount), Admin))
	http.Handle("/user-posts", Authorize(http.HandlerFunc(internal.HandleUserAccount), Admin))

	// reaction
	http.Handle("/reaction", http.HandlerFunc(internal.AllPostsHandler))

	// User pages handlers
	http.Handle("/edit", Authorize(http.HandlerFunc(internal.HandleUserEdit), Admin))
	http.Handle("/account", Authorize(http.HandlerFunc(internal.HandleUserAccount), Admin))

	// Auth
	http.Handle("/register", http.HandlerFunc(internal.HandleRegister))
	http.Handle("/login", http.HandlerFunc(internal.LoginHandler))
	http.Handle("/sign-out", Authorize(http.HandlerFunc(internal.HandleSignOut), Admin))

	// Oauth
	// http.Handle("/redirect", http.HandlerFunc(internal.Redirect_GET))
	http.Handle("/github/oauth/redirect", http.HandlerFunc(internal.HandleGithubRedirect))
	http.Handle("/github/oauth", http.HandlerFunc(internal.GithubAuthHandler))

	// http.Handle("/", Authorize(http.HandlerFunc(internal.ShowArtistsHandler), Admin))

	// static
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))

	// Custom handler

	// Run server
	port := getPort()
	fmt.Println("Server is listening...", "127.0.0.1"+port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	return ":" + port
}
