package auth

type UserRepository interface {
    GetByEmail(email string) (*User, error)
    Create(user *User) error
    UpdatePassword(email, passwordHash string) error
}
