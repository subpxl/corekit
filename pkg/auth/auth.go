package auth

// Main Auth service interface
type Auth interface {
    Register(email, password string) error
    Login(email, password string) (string, error)
    ForgotPassword(email string) error
    ResetPassword(token, newPassword string) error
}
