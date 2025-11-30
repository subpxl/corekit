package auth

import (
    "time"
)

type authService struct {
    repo          UserRepository
    tokenProvider TokenProvider
}

func NewAuthService(repo UserRepository, tokenProvider TokenProvider) Auth {
    return &authService{
        repo:          repo,
        tokenProvider: tokenProvider,
    }
}

func (s *authService) Register(email, password string) error {
    _, err := s.repo.GetByEmail(email)
    if err == nil {
        return ErrUserExists
    }

    hash, _ := HashPassword(password)
    user := &User{
        Email:        email,
        PasswordHash: hash,
        CreatedAt:    time.Now(),
    }

    return s.repo.Create(user)
}

func (s *authService) Login(email, password string) (string, error) {
    user, err := s.repo.GetByEmail(email)
    if err != nil {
        return "", ErrInvalidCredentials
    }

    if !CheckPasswordHash(password, user.PasswordHash) {
        return "", ErrInvalidCredentials
    }

    return s.tokenProvider.CreateToken(user.ID, user.Email)
}

func (s *authService) ForgotPassword(email string) error {
    user, err := s.repo.GetByEmail(email)
    if err != nil {
        return ErrNotFound
    }

    token, _ := s.tokenProvider.CreateResetToken(user.Email)
    // Here you would email token
    _ = token

    return nil
}

func (s *authService) ResetPassword(token, newPassword string) error {
    email, err := s.tokenProvider.VerifyResetToken(token)
    if err != nil {
        return err
    }

    newHash, _ := HashPassword(newPassword)
    return s.repo.UpdatePassword(email, newHash)
}
