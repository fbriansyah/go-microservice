package commands

import (
	"context"

	"github.com/fbriansyah/go-microservice/internal/application/domain/user"
	"github.com/fbriansyah/go-microservice/util"
)

type (
	CreateUserCmd struct {
		Name     string
		Email    string
		Password string
	}
	CreateUserHandler struct {
		userRepo user.Repository
	}
)

func NewUserHandler(userRepo user.Repository) CreateUserHandler {
	return CreateUserHandler{
		userRepo: userRepo,
	}
}

func (h CreateUserHandler) CreateUser(ctx context.Context, cmd CreateUserCmd) error {
	password, err := util.HashPassword(cmd.Password)
	if err != nil {
		return err
	}
	user := user.NewUser(cmd.Name, cmd.Email, password)
	return h.userRepo.Save(ctx, user)
}
