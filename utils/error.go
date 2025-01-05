package utils

import "errors"

var (
	ErrUserNotFound             = errors.New("username or password is not found")
	ErrUserNotAuthorized        = errors.New("user is not authorized")
	ErrTokenNotValid            = errors.New("token is not valid")
	ErrUsernamePasswordRequired = errors.New("username and password are required")
	ErrCreateUserRequiredParam  = errors.New("all param for create user are required")
)
