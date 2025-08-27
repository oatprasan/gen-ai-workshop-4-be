package infra

import (
	"gen-ai-workshop-4-be/internal/domain"
	"gen-ai-workshop-4-be/internal/ports"

	"gorm.io/gorm"
)

type GormUserRepo struct {
	db *gorm.DB
}

func NewGormUserRepo(db *gorm.DB) ports.UserRepository {
	return &GormUserRepo{db: db}
}

func (r *GormUserRepo) Create(u *domain.User) error {
	return r.db.Create(u).Error
}

func (r *GormUserRepo) FindByEmail(email string) (*domain.User, error) {
	var u domain.User
	if err := r.db.Where("email = ?", email).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *GormUserRepo) FindByID(id uint) (*domain.User, error) {
	var u domain.User
	if err := r.db.First(&u, id).Error; err != nil {
		return nil, err
	}
	return &u, nil
}
