package cmd

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	api "forum.com/api"
	config "forum.com/config"
	internal "forum.com/internal"
	model "forum.com/model"
)

func init() {
	config.InitConfig()
	model.Init()
}

const api_str = "/api/v1"

func Execute() {
	mux := http.NewServeMux()
	// http.
	// Api - V1
	mux.Handle("/log", http.HandlerFunc(internal.DriverLogHandler))
	// posts
	mux.Handle(api_str+"/posts", http.HandlerFunc(api.AllPostsHandler))
	mux.Handle(api_str+"/posts-by-category", http.HandlerFunc(api.AllPostsByCategoryHandler))
	mux.Handle(api_str+"/posts-by-user", Authorize(http.HandlerFunc(api.AllPostsByUserHandler), User))
	mux.Handle(api_str+"/posts-liked", Authorize(http.HandlerFunc(api.AllPostsLikedHandler), User))
	mux.Handle(api_str+"/like-post", Authorize(http.HandlerFunc(api.LikePostHandler), User))
	mux.Handle(api_str+"/dislike-post", Authorize(http.HandlerFunc(api.DislikePostHandler), User))
	// comments
	mux.Handle(api_str+"/comments", http.HandlerFunc(api.PostCommentsHandler))
	mux.Handle(api_str+"/comment/new", Authorize(http.HandlerFunc(api.NewCommentHandler), User))
	// Categories
	mux.Handle(api_str+"/categories", http.HandlerFunc(api.CategoryListHandler))

	mux.Handle("/", http.HandlerFunc(internal.Handle))
	// Posts
	mux.Handle("/post/new", Authorize(http.HandlerFunc(internal.NewPostHandler), User))
	mux.Handle("/post/edit", http.HandlerFunc(internal.NewEditHandler))
	mux.Handle("/post/delete", http.HandlerFunc(internal.NewEditHandler))
	mux.Handle("/posts", http.HandlerFunc(internal.AllPostsHandler))
	mux.Handle("/posts-category/", http.HandlerFunc(internal.AllPostsCategoryHandler))
	mux.Handle("/post/", http.HandlerFunc(internal.PostHandler))
	mux.Handle("/user-reacted-posts", Authorize(http.HandlerFunc(internal.UserReactedPostHandler), User))
	mux.Handle("/user-posts", Authorize(http.HandlerFunc(internal.HandleUserPosts), User))

	// reaction
	mux.Handle("/reaction", http.HandlerFunc(internal.AllPostsHandler))

	// User pages handlers
	mux.Handle("/edit", Authorize(http.HandlerFunc(internal.HandleUserEdit), User))
	mux.Handle("/account", Authorize(http.HandlerFunc(internal.HandleUserAccount), User))

	// Auth
	mux.Handle("/register", http.HandlerFunc(internal.HandleRegister))
	mux.Handle("/login", http.HandlerFunc(internal.LoginHandler))
	mux.Handle("/sign-out", Authorize(http.HandlerFunc(internal.HandleSignOut), User))

	// Oauth
	// mux.Handle("/redirect", http.HandlerFunc(internal.Redirect_GET))
	mux.Handle("/github/oauth/redirect", http.HandlerFunc(internal.HandleGithubRedirect))
	mux.Handle("/github/oauth", http.HandlerFunc(internal.GithubAuthHandler))

	// mux.Handle("/", Authorize(http.HandlerFunc(internal.ShowArtistsHandler), User))

	// static
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))

	// Custom handler

	// Run server
	app := App{port: getPort(), mux: mux}
	app.Run(mux)
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	return ":" + port
}

type App struct {
	port string
	mux  *http.ServeMux
}

func (a *App) Run(mux *http.ServeMux) {
	cer, err := tls.LoadX509KeyPair("config/crt/localhost.crt", "config/crt/localhost.key")
	if err != nil {
		log.Println(err)
		return
	}

	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  30 * time.Second,
		Addr:         ":https",
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{cer},
		},
		Handler:      limit(mux),
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}

	go func() {
		log.Fatal(http.ListenAndServe(":http", nil))
	}()

	fmt.Println("Server is listening...", "https://localhost")
	err = server.ListenAndServeTLS("", "")
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
