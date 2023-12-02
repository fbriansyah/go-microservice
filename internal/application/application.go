package application

import (
	"context"

	"github.com/fbriansyah/go-microservice/internal/application/commands"
	"github.com/fbriansyah/go-microservice/internal/application/domain/user"
	"github.com/fbriansyah/go-microservice/internal/application/queries"
	"github.com/rs/zerolog"
)

type (
	App interface {
		Commands
		Queries
	}
	Commands interface {
		CreateUser(ctx context.Context, cmd commands.CreateUserCmd) error
	}
	Queries interface {
		FindUserByEmail(ctx context.Context, query queries.FindUserByEmailQuery) (*user.User, error)
	}
	Application struct {
		userRepo user.Repository
		logger   zerolog.Logger
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

// WithUserRepo configures the application to use the specified user repository
func WithUserRepo(userRepo user.Repository) ApplicationConfig {
	return func(app *Application) error {
		app.userRepo = userRepo
		return nil
	}
}

func WithLogger(logger zerolog.Logger) ApplicationConfig {
	return func(app *Application) error {
		app.logger = logger
		return nil
	}
}

// New creates a new instance of the application
func New(cfgs ...ApplicationConfig) (*Application, error) {
	app := &Application{}

	// Iterate through the provided configuration functions,
	// calling each function in turn to configure the application
	for _, cfg := range cfgs {
		err := cfg(app)
		if err != nil {
			return nil, err
		}
	}

	// Registering the commands and queries
	app.appCommands = appCommands{
		CreateUserHandler: commands.NewUserHandler(app.userRepo, app.logger),
	}

	app.appQueries = appQueries{
		FindUserByEmailHandler: queries.NewFindUserByEmailHandler(app.userRepo),
	}

	return app, nil
}
