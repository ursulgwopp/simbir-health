package custom_errors

import "errors"

var (
	ErrUsernameExists            = errors.New("username already exists")
	ErrShortPassword             = errors.New("password must be at least 8 characters long")
	ErrPasswordWithoutDigits     = errors.New("password must contain at least one digit")
	ErrUsernameInvalidLength     = errors.New("username must be between 3 and 30 characters long")
	ErrUsernameInvalidCharacters = errors.New("username can only contain alphanumeric characters and underscores")
	ErrLastNameInvalid           = errors.New("last name must be between 2 and 30 characters long")
	ErrFirstNameInvalid          = errors.New("first name must be between 2 and 30 characters long")
	ErrSignIn                    = errors.New("invalid username or password")
	ErrInvalidToken              = errors.New("invalid token")
	ErrInvalidTokenType          = errors.New("token is of invalid type")
	ErrEmptyAuthHeader           = errors.New("empty auth header")
	ErrAccessDenied              = errors.New("access denied")
	ErrIdNotFound                = errors.New("id not found")
	ErrInvalidUserId             = errors.New("user id is of invalid type")
	ErrInvalidFrom               = errors.New("from can not be negative")
	ErrInvalidCount              = errors.New("count can not be negative")
	ErrInvalidName               = errors.New("invalid hospital name")
	ErrInvalidAddress            = errors.New("invalid hospital address")
	ErrInvalidPhone              = errors.New("invalid hospital contact phone")
)
