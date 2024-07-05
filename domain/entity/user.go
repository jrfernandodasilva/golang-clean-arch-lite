package entity

import (
	"errors"
	"net/mail"
)

var (
	ErrUserNotFound         = errors.New("User not found")
	ErrUserIdMustBePositive = errors.New("User id must be positive")
	ErrUserNameIsRequired   = errors.New("User name is required")
	ErrUserEmailIsRequired  = errors.New("User email is required")
)

type Role string

const (
	RoleAdmin Role = "admin"
	RoleGuest Role = "guest"
)

type User struct {
	ID    int    `faker:"oneof: 15, 27, 31, 55, 12, 93, 11, 112, 97, 21, 23, 57, 29"`
	Name  string `faker:"first_name"`
	Email string `faker:"email"`
	Role  Role   `faker:"-"`
}

func NewUser(id int, name string, email string, role Role) (*User, error) {
	user := &User{
		ID:    id,
		Name:  name,
		Email: email,
		Role:  role,
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) Validate() error {
	if u.ID <= 0 {
		return ErrUserIdMustBePositive
	}
	if u.Name == "" {
		return ErrUserNameIsRequired
	}
	if u.Email == "" {
		return ErrUserEmailIsRequired
	}
	if _, err := mail.ParseAddress(u.Email); err != nil {
		return err
	}

	return nil
}

func (u *User) DefaultRole() Role {
	if u.Role == "" {
		return RoleGuest
	}

	return u.Role
}
