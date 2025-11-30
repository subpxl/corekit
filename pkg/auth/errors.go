package auth

import "errors"

var (
    ErrUserExists         = errors.New("user already exists")
    ErrInvalidCredentials = errors.New("invalid email or password")
    ErrNotFound           = errors.New("user not found")
)
