	package auth

import "time"

// Minimal user model, safe to extend externally via embedding
type User struct {
    ID           int64
    Email        string
    PasswordHash string
    CreatedAt    time.Time
}