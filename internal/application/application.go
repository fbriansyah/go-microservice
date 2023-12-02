package application

import (
	"context"

	"github.com/fbriansyah/go-microservice/internal/application/commands"
	"github.com/fbriansyah/go-microservice/internal/application/domain/user"
	"github.com/fbriansyah/go-microservice/internal/application/queries"
)

type (
	App interface {
		Commands
	}
	Commands interface {
		CreateUser(ctx context.Context, cmd commands.CreateUserCmd) error
	}
	Queries interface {
		FindUserByEmail(ctx context.Context, query queries.FindUserByEmailQuery) (*user.User, error)
	}
	Application struct {
		userRepo user.Repository
		appCommands
		appQueries
	}
	appCommands struct {
		commands.CreateUserHandler
	}
	appQueries struct {
		queries.FindUserByEmailHandler
	}
)

var _ App = (*Application)(nil)

type ApplicationConfig func(*Application) error

func WithUserRepo(userRepo user.Repository) ApplicationConfig {
	return func(app *Application) error {
		app.userRepo = userRepo
		return nil
	}
}

func New(cfgs ...ApplicationConfig) (*Application, error) {
	app := &Application{}

	for _, cfg := range cfgs {
		err := cfg(app)
		if err != nil {
			return nil, err
		}
	}

	app.appCommands = appCommands{
		CreateUserHandler: commands.NewUserHandler(app.userRepo),
	}

	app.appQueries = appQueries{
		FindUserByEmailHandler: queries.NewFindUserByEmailHandler(app.userRepo),
	}

	return app, nil
}
