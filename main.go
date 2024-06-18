package main

import (
	"context"
	"database/sql"
	"github.com/heroiclabs/nakama-common/runtime"
	_ "github.com/lib/pq"
	"hash-file-service/database"
	"hash-file-service/service"
)

func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	logger.Info("Running migrations...")
	if err := database.RunMigrations(db, logger); err != nil {
		logger.Error("Failed to run migrations", "error", err)
		return err
	}

	if err := service.RegisterFileProcessingRpc(logger, initializer); err != nil {
		logger.Error("Unable to register file processing RPC", "error", err)
		return err
	}

	return nil
}
