package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"gen-ai-workshop-4-be/internal/domain"
	"gen-ai-workshop-4-be/internal/service"
)

type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) Create(u *domain.User) error {
	args := m.Called(u)
	return args.Error(0)
}
func (m *MockUserRepo) FindByEmail(email string) (*domain.User, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}
func (m *MockUserRepo) FindByID(id uint) (*domain.User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

func TestRegister_Success(t *testing.T) {
	repo := new(MockUserRepo)
	svc := service.NewAuthService(repo)

	u := &domain.User{Email: "a@b.com", FirstName: "John", LastName: "Doe", Phone: "012"}

	repo.On("FindByEmail", "a@b.com").Return(nil, nil)
	repo.On("Create", mock.Anything).Run(func(args mock.Arguments) {
		u := args.Get(0).(*domain.User)
		u.ID = 1
	}).Return(nil)

	err := svc.Register(u, "secret123")
	assert.NoError(t, err)
	assert.Equal(t, uint(1), u.ID)
	// password should be hashed
	assert.NotEqual(t, "secret123", u.Password)
}

func TestAuthenticate_Success(t *testing.T) {
	repo := new(MockUserRepo)
	svc := service.NewAuthService(repo)

	// prepare hashed password
	u := &domain.User{Email: "a@b.com", FirstName: "John"}
	hash, _ := service.HashPassword("secret123")
	u.Password = hash

	repo.On("FindByEmail", "a@b.com").Return(u, nil)

	res, err := svc.Authenticate("a@b.com", "secret123")
	assert.NoError(t, err)
	assert.Equal(t, u.Email, res.Email)
}
