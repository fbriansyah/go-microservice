package main

import (
	"database/sql"
	"math"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func (m *Module) runDBMigration() {
	migration, err := migrate.New(m.config.MigrationURL, m.config.DBSource)
	if err != nil {
		m.logger.Fatal().Err(err).Msg("cannot create new migrate instance")
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		m.logger.Fatal().Err(err).Msg("failed to run migrate up")
	}

	m.logger.Info().Msg("db migrated successfully")
}

func (m *Module) openDB() (*sql.DB, error) {
	db, err := sql.Open(m.config.DBDriver, m.config.DBSource)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (m *Module) connectToDB() *sql.DB {
	var connectionCounts int64
	var backOff = 1 * time.Second
	for {
		connection, err := m.openDB()
		if err != nil {
			m.logger.Info().Msg("Postgres not yet ready ...")
			connectionCounts++
		} else {
			m.logger.Info().Msg("Connected to Postgres!")
			return connection
		}

		if connectionCounts > 10 {
			m.logger.Fatal().Err(err)
			return nil
		}

		backOff = time.Duration(math.Pow(float64(connectionCounts), 2)) * time.Second
		m.logger.Info().Msg("backing off...")
		time.Sleep(backOff)
		continue
	}
}
