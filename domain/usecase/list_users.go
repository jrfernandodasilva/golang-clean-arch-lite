package usecase

import (
	"github.com/jrfernandodasilva/golang-clean-arch-lite/domain/repository"
)

type UserDTO struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type ListUserOutputDTO struct {
	Users []UserDTO `json:"users"`
}

type ListUsersUseCase struct {
	repo repository.UserRepository
}

func NewListUsersUseCase(repo repository.UserRepository) *ListUsersUseCase {
	return &ListUsersUseCase{repo: repo}
}

func (uc *ListUsersUseCase) Execute() (*ListUserOutputDTO, error) {
	users, err := uc.repo.ListUsers()
	if err != nil {
		return nil, err
	}

	userDTOs := make([]UserDTO, len(users))
	for i, user := range users {
		userDTOs[i] = UserDTO{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Role:  string(user.Role),
		}
	}

	return &ListUserOutputDTO{Users: userDTOs}, nil
}
