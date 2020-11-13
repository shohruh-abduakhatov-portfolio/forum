package model

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Session struct {
	ID     string
	UserID string
	Expiry time.Time
}

const (
	// Keep users logged in for 3 Hours
	sessionLength     = 3 * time.Hour
	sessionCookieName = "GophrSession"
	sessionIDLength   = 20
)

func NewSession(w http.ResponseWriter) *Session {
	timeout := time.Hour * 1
	if w.Header().Get("timeout") != "" {
		timeout *= 24 * 365 * 10
	}
	expiry := time.Now().Add(sessionLength).Add(timeout)

	u2, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return nil
	}

	session := &Session{
		ID:     u2.String(),
		Expiry: expiry,
	}

	cookie := http.Cookie{
		Name:    sessionCookieName,
		Value:   session.ID,
		Expires: expiry,
		Path:    "/",
	}

	http.SetCookie(w, &cookie)
	return session

}

func RequestSession(r *http.Request) *Session {
	cookie, err := r.Cookie(sessionCookieName)
	if err != nil {
		return nil
	}
	session, err := GlobalSessionStore.Find(cookie.Value)
	if err != nil && err != ErrNotFound {
		panic(err)
	}
	if session == nil {
		return nil
	}
	if session.Expired() {
		GlobalSessionStore.Delete(session)
		return nil
	}
	return session
}

func (session *Session) Expired() bool {
	return session.Expiry.Before(time.Now())
}

func RequestUser(r *http.Request) *User {
	session := RequestSession(r)
	if session == nil || session.UserID == "" {
		return nil
	}
	user, err := GlobalUserStore.Find(session.UserID)
	if err != nil {
		panic(err)
	}
	return user
}

func RequireLogin(w http.ResponseWriter, r *http.Request) {
	// Let the request pass if we've got a user
	if RequestUser(r) != nil {
		return
	}
	query := url.Values{}
	query.Add("next", url.QueryEscape(r.URL.String()))
	http.Redirect(w, r, "/login?"+query.Encode(), http.StatusFound)
}

func FindOrCreateSession(w http.ResponseWriter, r *http.Request) *Session {
	session := RequestSession(r)
	if session == nil {
		session = NewSession(w)
	}
	return session
}
