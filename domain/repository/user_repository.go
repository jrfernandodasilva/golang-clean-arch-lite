package repository

import "github.com/jrfernandodasilva/golang-clean-arch-lite/domain/entity"

type UserRepository interface {
	CreateUser(user *entity.User) error
	FindUserByID(id int) (*entity.User, error)
	ListUsers() ([]entity.User, error)
}
