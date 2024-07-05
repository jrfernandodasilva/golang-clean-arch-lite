package persistence

import (
	"github.com/jrfernandodasilva/golang-clean-arch-lite/domain/entity"
	"github.com/jrfernandodasilva/golang-clean-arch-lite/domain/repository"
)

type MemoryUserRepository struct {
	users map[int]*entity.User
}

func NewMemoryUserRepository() repository.UserRepository {
	return &MemoryUserRepository{
		users: make(map[int]*entity.User),
	}
}

func (r *MemoryUserRepository) CreateUser(user *entity.User) error {
	r.users[user.ID] = user
	return nil
}

func (r *MemoryUserRepository) FindUserByID(id int) (*entity.User, error) {
	user, ok := r.users[id]
	if !ok {
		return nil, entity.ErrUserNotFound
	}
	return user, nil
}

func (r *MemoryUserRepository) ListUsers() ([]entity.User, error) {

	users := make([]entity.User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, *user)
	}
	return users, nil
}
