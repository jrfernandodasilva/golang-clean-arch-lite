package usecase

import (
	"github.com/jrfernandodasilva/golang-clean-arch-lite/domain/repository"
)

type GetUserInputDTO struct {
	ID int
}

type GetUserOutputDTO struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type GetUserUserCase struct {
	repo repository.UserRepository
}

func NewGetUser(repo repository.UserRepository) *GetUserUserCase {
	return &GetUserUserCase{repo: repo}
}

func (uc *GetUserUserCase) Execute(input GetUserInputDTO) (*GetUserOutputDTO, error) {
	user, err := uc.repo.FindUserByID(input.ID)
	if err != nil {
		return nil, err
	}

	return &GetUserOutputDTO{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  string(user.Role),
	}, nil
}
