package service

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"

	"gen-ai-workshop-4-be/internal/domain"
	"gen-ai-workshop-4-be/internal/ports"
)

var ErrUserExists = errors.New("user already exists")

// AuthService contains business logic for auth
type AuthService struct {
	repo ports.UserRepository
}

func NewAuthService(r ports.UserRepository) *AuthService {
	return &AuthService{repo: r}
}

func (s *AuthService) Repo() ports.UserRepository {
	return s.repo
}

func HashPassword(plain string) (string, error) {
	h, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(h), err
}

func (s *AuthService) Register(u *domain.User, plainPassword string) error {
	// check existing
	if ex, _ := s.repo.FindByEmail(u.Email); ex != nil {
		return ErrUserExists
	}

	// hash
	hash, err := HashPassword(plainPassword)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	u.CreatedAt = time.Now()

	return s.repo.Create(u)
}

func (s *AuthService) Authenticate(email, plain string) (*domain.User, error) {
	u, err := s.repo.FindByEmail(email)
	if err != nil || u == nil {
		return nil, errors.New("invalid credentials")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plain)); err != nil {
		return nil, errors.New("invalid credentials")
	}
	return u, nil
}
