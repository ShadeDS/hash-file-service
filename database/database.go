package database

import (
	"database/sql"
	"errors"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/heroiclabs/nakama-common/runtime"
)

// RunMigrations applies the database migrations.
func RunMigrations(db *sql.DB, logger runtime.Logger) error {
	logger.Info("Starting database migrations...")
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		logger.Error("Failed to create migration driver", "error", err)
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:///nakama/data/migrations",
		"postgres", driver)
	if err != nil {
		logger.Error("Failed to create migration instance", "error", err)
		return err
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		logger.Error("Failed to apply migrations", "error", err)
		return err
	}

	logger.Info("Database migrations applied successfully")
	return nil
}

// InsertData inserts file metadata into the database.
func InsertData(db *sql.DB, logger runtime.Logger, fileType, version, hash string, content []byte) error {
	logger.Info("Inserting data into the database", "type", fileType, "version", version, "hash", hash)
	query := `insert into files (type, version, content, hash) values ($1, $2, $3, $4)`
	_, err := db.Exec(query, fileType, version, content, hash)
	if err != nil {
		logger.Error("Failed to insert data", "error", err)
		return err
	}
	logger.Info("Data inserted successfully")
	return nil
}
