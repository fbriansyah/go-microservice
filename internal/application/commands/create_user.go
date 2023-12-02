package commands

import (
	"context"
	"errors"

	"github.com/fbriansyah/go-microservice/internal/application/domain/user"
	"github.com/fbriansyah/go-microservice/util"
	"github.com/rs/zerolog"
)

type (
	CreateUserCmd struct {
		Name     string
		Email    string
		Password string
	}
	CreateUserHandler struct {
		userRepo user.Repository
		logger   zerolog.Logger
	}
)

func NewUserHandler(userRepo user.Repository, logger zerolog.Logger) CreateUserHandler {
	return CreateUserHandler{
		userRepo: userRepo,
		logger:   logger,
	}
}

func (h CreateUserHandler) CreateUser(ctx context.Context, cmd CreateUserCmd) error {
	password, err := util.HashPassword(cmd.Password)
	if err != nil {
		h.logger.Error().Err(err).Msg("failed to hash password")
		return err
	}
	user := user.NewUser(cmd.Name, cmd.Email, password)
	err = h.userRepo.Save(ctx, user)
	if err != nil {
		h.logger.Error().Err(err).Msgf("failed to save user: %s", err.Error())
		return errors.New("failed to save user")
	}
	return nil
}
