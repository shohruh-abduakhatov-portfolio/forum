package model

import (
	"fmt"
	"time"

	// sqlite "./sqlite"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           string    `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	RoleID       int       `json:"role_id"`
	PermissionID int       `json:"permission_id"`
	PhotoID      int       `json:"photo_id"`
	DateCreated  time.Time `json:"date_created"`
}

const (
	passwordLength = 8
	hashCost       = 10
	userIDLength   = 16
)

func NewUser(username, email, password string) (User, error) {
	user := User{
		Email:    email,
		Username: username,
	}
	if username == "" {
		return user, errNoUsername
	}
	if email == "" {
		return user, errNoEmail
	}
	if password == "" {
		return user, errNoPassword
	}
	if len(password) < passwordLength {
		return user, errPasswordTooShort
	}

	// Check if the username exists
	// existingUser, err := GlobalUserStore.FindByUsername(username)
	existingUser, err := GlobalUserStore.FindByUsername(username)

	// DB
	if err != nil {
		if err != ErrNotFound {
			return user, err
		}
	}
	if existingUser != nil {
		return user, errUsernameExists
	}

	// Check if the email exists
	existingUser, err = GlobalUserStore.FindByEmail(email)
	if err != nil {
		if err != ErrNotFound {
			return user, err
		}
	}
	if existingUser != nil {
		return user, errEmailExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), hashCost)
	user.Password = string(hashedPassword)

	u2, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return user, err
	}
	user.ID = u2.String()
	return user, err
}

func t() {
	u1 := uuid.Must(uuid.NewV4())
	fmt.Printf("UUIDv4: %s\n", u1)

	// or error handling
	u2, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return
	}
	fmt.Printf("UUIDv4: %s\n", u2)

	// Parsing UUID from string input
	u2, err = uuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return
	}
	fmt.Printf("Successfully parsed: %s", u2)
}

func FindUser(username, password string) (*User, error) {
	out := &User{
		Username: username,
	}
	existingUser, err := GlobalUserStore.FindByUsername(username)
	if err != nil {
		return out, err
	}
	if existingUser == nil {
		return out, errCredentialsIncorrect
	}
	if bcrypt.CompareHashAndPassword(
		[]byte(existingUser.Password),
		[]byte(password),
	) != nil {
		return out, errCredentialsIncorrect
	}
	return existingUser, nil
}

func UpdateUser(user *User, email, currentPassword, newPassword string) (User, error) {
	out := *user
	out.Email = email
	// Check if the email exists
	existingUser, err := GlobalUserStore.FindByEmail(email)
	if err != nil {
		return out, err
	}
	if existingUser != nil && existingUser.ID != user.ID {
		return out, errEmailExists
	}
	// At this point, we can update the email address
	user.Email = email
	// No current password? Don't try update the password.
	if currentPassword == "" {
		return out, nil
	}
	if bcrypt.CompareHashAndPassword(
		[]byte(user.Password), []byte(currentPassword),
	) != nil {
		return out, errPasswordIncorrect
	}
	if newPassword == "" {
		return out, errNoPassword
	}
	if len(newPassword) < passwordLength {
		return out, errPasswordTooShort
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), hashCost)
	user.Password = string(hashedPassword)
	return out, err
}
