package model

import (
	"errors"

	"github.com/mattn/go-sqlite3"
)

type ValidationError error
type ItemError error

var (
	errNoUsername       = ValidationError(errors.New("You must supply a username"))
	errNoEmail          = ValidationError(errors.New("You must supply an email"))
	errNoPassword       = ValidationError(errors.New("You must supply a password"))
	errPasswordTooShort = ValidationError(errors.New("Your passwordis too short"))
	errUsernameExists   = ValidationError(errors.New("That username is taken"))
	errEmailExists      = ValidationError(errors.New("That email address has an account"))

	errInvalidCommentString = ValidationError(errors.New("Invalid comment string"))
	errCommentContentMinLen = ValidationError(errors.New("Invalid comment content min len"))

	errInvalidTopicString = ValidationError(errors.New("Invalid topic string"))
	errTopicContentMinLen = ValidationError(errors.New("Invalid topic content min len"))
	errNoTopicFound       = ValidationError(errors.New("No such Post found"))

	errTextContentMinLen = ValidationError(errors.New("Invalid topic content min len"))
	errImage             = ValidationError(errors.New("Invalid Image"))

	errSomethingWentWrong = ValidationError(errors.New("Something went wrong"))
)

var (
	// ErrNotFound means the requested item is not found.
	ErrNotFound = ItemError(errors.New("store: item not found"))
	// ErrConflict means the operation failed because of a conflict between items.
	ErrConflict = ItemError(errors.New("store: item conflict"))
	// ErrNotInserted means the operation failed because no new item inserted.
	ErrNotInserted = ItemError(errors.New("store: not inserted"))

	ErrImageFormat          = ValidationError(errors.New("Invalid Image format"))
	errCredentialsIncorrect = ItemError(errors.New("Incorrect username/password"))
	errPasswordIncorrect    = ItemError(errors.New("Password did not match"))
)

func ErrWentWrong() ValidationError {
	return errSomethingWentWrong
}

func IsValidationError(err error) bool {
	_, ok := err.(ValidationError)
	return ok
}

func IsItemError(err error) bool {
	_, ok := err.(ItemError)
	return ok
}

func IsSqliteError(err error) bool {
	// err, ok := err.(*go-sqlite3.Error); ok
	_, ok := err.(sqlite3.Error)
	return ok
}

func IsUniqueConstraintError(err error) bool {
	if err == sqlite3.ErrConstraintUnique {
		return true
	}
	return false
}
