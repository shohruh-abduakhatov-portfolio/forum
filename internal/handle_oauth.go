package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	model "forum.com/model"
	render "forum.com/render"
)

func Redirect_GET(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	render.Basic(w, req, "redirect.html", map[string]interface{}{})
}

func GithubAuthHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/github/oauth" {
		http.Error(w, "Go back to the main page", 404)
		return
	}
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(os.Stdout, "could not parse query: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	token := r.FormValue("token")
	ghu, err := model.NewGithubUser(
		model.OAuthAccessResponse{
			AccessToken: token,
		})
	if err != nil {
		fmt.Fprintf(os.Stdout, "could not send HTTP request: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	// Create User object
	user, err := model.NewUser(
		ghu.Username,
		ghu.Email,
		ghu.Password,
	)
	if err != nil {
		if model.IsValidationError(err) {
			// Means email exists
			tmpUser, err := model.GlobalUserStore.FindByEmail(user.Email)
			if err != nil {
				http.Redirect(w, r, "/login?flash=Something+went+wrong.", http.StatusInternalServerError)
				return
			}
			user = *tmpUser
		} else {
			// Unexpected error
			http.Redirect(w, r, "/login?flash=Something+went+wrong.", http.StatusInternalServerError)
			return
		}

	} else {
		// save
		_, err = model.GlobalUserStore.New(user)
		if err != nil {
			http.Redirect(w, r, "/login?flash=Something+went+wrong.", http.StatusFound)
			return
		}
	}

	// Create a new session
	session := model.FindOrCreateSession(w, r)
	// cookie, err := r.Cookie("GophrSession")
	// if err != nil {
	// 	return
	// }
	// fmt.Println("cookie", cookie)

	session.UserID = user.ID
	err = model.GlobalSessionStore.Save(session)
	if err != nil {
		if !model.IsSqliteError(err) && !model.IsUniqueConstraintError(err) {
			panic(err)
		}

	}
	// render.Template(w, r, "post/posts", map[string]interface{}{})
	http.Redirect(w, r, "/posts?flash=Succesfully+authorized+via+Github!", http.StatusFound)
}

func HandleGithubRedirect(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/github/oauth/redirect" {
		http.Error(w, "Go back to the main page", 404)
		return
	}
	// We will be using `httpClient` to make external HTTP requests later in our code
	httpClient := http.Client{}
	clientID := os.Getenv("GithubClientID")
	clientSecret := os.Getenv("GithubSecretKey")

	// First, we need to get the value of the `code` query param
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(os.Stdout, "could not parse query: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	code := r.FormValue("code")

	// Next, lets for the HTTP request to call the github oauth enpoint
	// to get our access token
	reqURL := fmt.Sprintf("https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s", clientID, clientSecret, code)
	req, err := http.NewRequest(http.MethodPost, reqURL, nil)
	if err != nil {
		fmt.Fprintf(os.Stdout, "could not create HTTP request: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	// We set this header since we want the response
	// as JSON
	req.Header.Set("accept", "application/json")

	// Send out the HTTP request
	res, err := httpClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stdout, "could not send HTTP request: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer res.Body.Close()

	var t = model.OAuthAccessResponse{}
	// Parse the request body into the `OAuthAccessResponse` struct
	if err := json.NewDecoder(res.Body).Decode(&t); err != nil {
		fmt.Fprintf(os.Stdout, "could not parse JSON response: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	// cookie, err := r.Cookie("GophrSession")
	// if err == nil && cookie != nil {
	// 	fmt.Println("cookie", cookie)
	// }

	// http.SetCookie(w, nil)
	http.Redirect(w, r, "/github/oauth?token="+t.AccessToken, http.StatusFound)

	// w.Header().Set("Location", "/github/oauth?from=github&access_token="+t.AccessToken)
	// w.WriteHeader(http.StatusFound)
}
