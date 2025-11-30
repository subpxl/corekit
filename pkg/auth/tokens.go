package auth

// token provider interface enables JWT, session tokens, etc
type TokenProvider interface {
	CreateToken(userID int64, email string) (string, error)
	CreateResetToken(email string) (string, error)
	VerifyResetToken(token string) (string, error)
}
