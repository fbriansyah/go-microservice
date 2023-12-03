package application

import (
	"database/sql"
	"testing"

	"github.com/fbriansyah/go-microservice/internal/application/domain/user"
	"github.com/fbriansyah/go-microservice/internal/logging"
	"github.com/fbriansyah/go-microservice/internal/postgres"
	"github.com/fbriansyah/go-microservice/util"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var app *Application
var logger zerolog.Logger
var userRepo user.Repository

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	logger = logging.New(logging.LogConfig{
		Environment: config.Environment,
		LogLevel:    logging.DEBUG,
	})

	db, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		logger.Fatal().Msg(err.Error())
	}

	userRepo = postgres.NewUserRepo(db, "users.users")

	app, err = New(
		WithUserRepo(userRepo),
		WithLogger(logger),
	)

	if err != nil {
		logger.Fatal().Msg(err.Error())
	}

	m.Run()
}
