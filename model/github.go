package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

type OAuthAccessResponse struct {
	AccessToken string `json:"access_token"`
}

type GithubUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"Username"`
}

func NewGithubUser(token OAuthAccessResponse) (GithubUser, error) {
	user, err := fetchGithubUserData(token.AccessToken)
	if err != nil {
		return user, err
	}
	// Random password
	u1, err := uuid.NewV1()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return user, err
	}
	user.Password = u1.String()

	// Random username
	u1, err = uuid.NewV1()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return user, err
	}
	user.Username = u1.String()
	return user, nil
}

func fetchGithubUserData(token string) (user GithubUser, err error) {
	url := "https://api.github.com/user"
	client := &http.Client{}
	defer client.CloseIdleConnections()
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "token "+token)
	if err != nil {
		return
	}
	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	githubUser := GithubUser{}
	jsonErr := json.Unmarshal(body, &githubUser)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return githubUser, nil
}
