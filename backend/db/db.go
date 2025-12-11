package db

import (
	"database/sql"
	"embed"
	"os"

	_ "github.com/lib/pq"

	"github.com/elect0/xpose/backend/internal/platform/logger"
	migrate "github.com/rubenv/sql-migrate"
	"go.uber.org/zap"
)

//go:embed migrations/*
var migrationsFS embed.FS

func New(uri, environment string, logger *logger.Logger) (*sql.DB, error) {
	db, err := sql.Open("postgres", uri)
	if err != nil {
		return nil, err
	}

	migrations := &migrate.EmbedFileSystemMigrationSource{
		FileSystem: migrationsFS,
		Root:       "migrations",
	}

	if _, err := migrate.Exec(db, "postgres", migrations, migrate.Up); err != nil {
		return nil, err
	}

	if environment == "development" {
		seedFiles, err := os.ReadDir("db/seed")
		if err != nil {
			logger.Error("Failed to read seed files", zap.Error(err))
		}

		for _, file := range seedFiles {
			c, err := os.ReadFile("db/seed" + file.Name())
			if err != nil {
				logger.Error("Error while reading seed file", zap.Error(err))
			}

			sqlCode := string(c)

			_, err = db.Exec(sqlCode)
			if err != nil {
				logger.Error("Failed to seed database", zap.Error(err))
			}
		}

		logger.Info("Seeded database")
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
