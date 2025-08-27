package ports

import "gen-ai-workshop-4-be/internal/domain"

// UserRepository defines persistence operations used by services
type UserRepository interface {
	Create(u *domain.User) error
	FindByEmail(email string) (*domain.User, error)
	FindByID(id uint) (*domain.User, error)
}
