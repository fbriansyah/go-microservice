package application

import (
	"context"

	"github.com/fbriansyah/go-microservice/internal/application/commands"
	"github.com/fbriansyah/go-microservice/internal/application/domain/user"
)

type (
	App interface {
		Commands
	}
	Commands interface {
		CreateUser(ctx context.Context, cmd *commands.CreateUserCmd) error
	}
	Application struct {
		userRepo user.Repository
		appCommands
	}
	appCommands struct {
		commands.CreateUserHandler
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

	return app, nil
}
