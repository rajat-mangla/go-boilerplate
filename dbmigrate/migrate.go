package dbmigrate

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rajat-mangla/go-boilerplate/config"
	"github.com/rajat-mangla/go-boilerplate/repository"
)

// Up applies all pending migrations from the migrations directory.
func Up(cfg config.DatabaseConfig) error {
	return run(cfg, func(m *migrate.Migrate) error {
		return m.Up()
	})
}

// Down rolls back the most recently applied migration.
func Down(cfg config.DatabaseConfig) error {
	return run(cfg, func(m *migrate.Migrate) error {
		return m.Steps(-1)
	})
}

func run(cfg config.DatabaseConfig, apply func(*migrate.Migrate) error) error {
	databaseURL, err := repository.ConnectionURL(cfg)
	if err != nil {
		return err
	}

	migrationsPath, err := filepath.Abs("migrations")
	if err != nil {
		return fmt.Errorf("resolve migrations directory: %w", err)
	}
	migrationsURL := (&url.URL{
		Scheme: "file",
		Path:   filepath.ToSlash(migrationsPath),
	}).String()

	m, err := migrate.New(migrationsURL, databaseURL)
	if err != nil {
		return fmt.Errorf("initialize migration runner: %w", err)
	}

	operationErr := apply(m)
	if errors.Is(operationErr, migrate.ErrNoChange) {
		operationErr = nil
	}

	sourceErr, databaseErr := m.Close()
	return errors.Join(operationErr, sourceErr, databaseErr)
}

// Create writes a pair of up/down SQL migration files prefixed by Unix epoch seconds.
func Create(directory, name string) (string, string, error) {
	if err := os.MkdirAll(directory, 0755); err != nil {
		return "", "", fmt.Errorf("create migrations directory: %w", err)
	}

	version := strconv.FormatInt(time.Now().UTC().Unix(), 10)
	base := filepath.Join(directory, version+"_"+name)
	upPath := base + ".up.sql"
	downPath := base + ".down.sql"
	if err := createPair(upPath, downPath); err != nil {
		return "", "", err
	}
	return upPath, downPath, nil
}

func createPair(upPath, downPath string) error {
	if err := createMigrationFile(upPath, "-- Add schema changes in this migration.\n"); err != nil {
		return fmt.Errorf("create up migration: %w", err)
	}

	if err := createMigrationFile(downPath, "-- Revert schema changes in this migration.\n"); err != nil {
		return errors.Join(fmt.Errorf("create down migration: %w", err), removeMigrationFiles(upPath))
	}
	return nil
}

func createMigrationFile(path, contents string) error {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		return err
	}

	_, writeErr := file.WriteString(contents)
	closeErr := file.Close()
	if err := errors.Join(writeErr, closeErr); err != nil {
		return errors.Join(err, removeMigrationFiles(path))
	}
	return nil
}

func removeMigrationFiles(paths ...string) error {
	var errs []error
	for _, path := range paths {
		if err := os.Remove(path); err != nil && !errors.Is(err, os.ErrNotExist) {
			errs = append(errs, fmt.Errorf("remove incomplete migration %q: %w", path, err))
		}
	}
	return errors.Join(errs...)
}
