package queries

import (
	"context"
	"errors"

	"github.com/fbriansyah/go-microservice/internal/application/domain/user"
	"github.com/fbriansyah/go-microservice/util"
)

var (
	ErrCannotFindUser  = errors.New("cannot find user")
	ErrInvalidPassword = errors.New("invalid password")
)

type (
	FindUserByEmailQuery struct {
		Email    string
		Password string
	}
	FindUserByEmailHandler struct {
		userRepo user.Repository
	}
)

func NewFindUserByEmailHandler(userRepo user.Repository) FindUserByEmailHandler {
	return FindUserByEmailHandler{
		userRepo: userRepo,
	}
}

func (h FindUserByEmailHandler) FindUserByEmail(ctx context.Context, query FindUserByEmailQuery) (*user.User, error) {
	// TODO: Add session id to check if the user is already logged in.
	u, err := h.userRepo.FindByEmail(ctx, query.Email)
	if err != nil {
		return &user.User{}, ErrCannotFindUser
	}
	if err := util.CheckPassword(query.Password, u.Password); err != nil {
		return &user.User{}, ErrInvalidPassword
	}

	return u, nil
}
