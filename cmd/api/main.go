package main

import (
	"github.com/fbriansyah/go-microservice/internal/application"
	"github.com/fbriansyah/go-microservice/internal/echo"
	"github.com/fbriansyah/go-microservice/internal/logging"
	"github.com/fbriansyah/go-microservice/internal/postgres"
	"github.com/fbriansyah/go-microservice/util"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Module struct {
	config *util.Config
	logger zerolog.Logger
}

func NewModule(config *util.Config) *Module {
	logger := logging.New(logging.LogConfig{
		Environment: config.Environment,
		LogLevel:    logging.DEBUG,
	})

	return &Module{
		config: config,
		logger: logger,
	}
}

func main() {
	config, err := util.LoadConfig("./")
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	module := NewModule(&config)

	postgresDB := module.connectToDB()
	if postgresDB == nil {
		module.logger.Fatal().Msgf("cannot connect to db: %s", err.Error())
	}

	module.runDBMigration()

	userRepo := postgres.NewUserRepo(postgresDB, "users.users")

	app, err := application.New(
		application.WithUserRepo(userRepo),
	)
	if err != nil {
		module.logger.Fatal().Msg(err.Error())
	}
	server := echo.NewServer(app, config.RestURL, module.logger)
	server.Start()
}
