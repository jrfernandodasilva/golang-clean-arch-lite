package usecase

import (
	"github.com/jrfernandodasilva/golang-clean-arch-lite/domain/entity"
	"github.com/jrfernandodasilva/golang-clean-arch-lite/domain/repository"
)

type CreateUserInputDTO struct {
	ID    int
	Name  string
	Email string
}

type CreateUserOutputDTO struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type CreateUserUseCase struct {
	repo repository.UserRepository
}

func NewCreateUser(repo repository.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{repo: repo}
}

func (uc *CreateUserUseCase) Execute(input CreateUserInputDTO) (CreateUserOutputDTO, error) {
	user, err := entity.NewUser(
		input.ID,
		input.Name,
		input.Email,
		entity.RoleGuest,
	)

	if err != nil {
		return CreateUserOutputDTO{}, err
	}

	err = uc.repo.CreateUser(user)
	if err != nil {
		return CreateUserOutputDTO{}, err
	}

	output := CreateUserOutputDTO{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  string(user.Role),
	}

	return output, nil
}
