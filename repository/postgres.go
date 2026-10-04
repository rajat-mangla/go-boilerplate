package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net"
	"net/url"
	"strconv"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/rajat-mangla/go-boilerplate/config"
)

// Repository provides database/sql operations for PostgreSQL.
type Repository struct {
	db *sql.DB
}

// Open creates a PostgreSQL connection pool and verifies connectivity.
func Open(ctx context.Context, cfg config.DatabaseConfig) (*Repository, error) {
	if cfg.MaxOpenConns < 0 {
		return nil, errors.New("postgres max open connections cannot be negative")
	}
	if cfg.MaxIdleConns < 0 {
		return nil, errors.New("postgres max idle connections cannot be negative")
	}
	if cfg.MaxOpenConns > 0 && cfg.MaxIdleConns > cfg.MaxOpenConns {
		return nil, errors.New("postgres max idle connections cannot exceed max open connections")
	}
	if cfg.ConnectTimeout < 0 || cfg.ConnMaxLifetime < 0 || cfg.ConnMaxIdleTime < 0 {
		return nil, errors.New("postgres connection timeouts and lifetimes cannot be negative")
	}

	connectionURL, err := ConnectionURL(cfg)
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("pgx", connectionURL)
	if err != nil {
		return nil, fmt.Errorf("open postgres connection: %w", err)
	}

	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetConnMaxLifetime(cfg.ConnMaxLifetime)
	db.SetConnMaxIdleTime(cfg.ConnMaxIdleTime)

	pingCtx := ctx
	cancel := func() {}
	if cfg.ConnectTimeout > 0 {
		pingCtx, cancel = context.WithTimeout(ctx, cfg.ConnectTimeout)
	}
	defer cancel()

	if err := db.PingContext(pingCtx); err != nil {
		closeErr := db.Close()
		if closeErr != nil {
			return nil, errors.Join(
				fmt.Errorf("ping postgres: %w", err),
				fmt.Errorf("close postgres connection: %w", closeErr),
			)
		}
		return nil, fmt.Errorf("ping postgres: %w", err)
	}

	return &Repository{db: db}, nil
}

// ConnectionURL builds a PostgreSQL URL from the configured connection fields.
func ConnectionURL(cfg config.DatabaseConfig) (string, error) {
	if cfg.Host == "" {
		return "", errors.New("postgres host is required")
	}
	if cfg.Port < 1 || cfg.Port > 65535 {
		return "", errors.New("postgres port must be between 1 and 65535")
	}
	if cfg.User == "" {
		return "", errors.New("postgres user is required")
	}
	if cfg.Name == "" {
		return "", errors.New("postgres database name is required")
	}

	query := url.Values{"sslmode": []string{"disable"}}
	return (&url.URL{
		Scheme:   "postgres",
		Host:     net.JoinHostPort(cfg.Host, strconv.Itoa(cfg.Port)),
		User:     url.UserPassword(cfg.User, cfg.Password),
		Path:     "/" + cfg.Name,
		RawQuery: query.Encode(),
	}).String(), nil
}

// ExecContext executes a query that does not return rows.
func (r *Repository) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return r.db.ExecContext(ctx, query, args...)
}

// QueryContext executes a query that returns rows.
func (r *Repository) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	return r.db.QueryContext(ctx, query, args...)
}

// QueryRowContext executes a query expected to return at most one row.
func (r *Repository) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	return r.db.QueryRowContext(ctx, query, args...)
}

// BeginTx starts a transaction.
func (r *Repository) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	return r.db.BeginTx(ctx, opts)
}

// PingContext verifies that the database is reachable.
func (r *Repository) PingContext(ctx context.Context) error {
	return r.db.PingContext(ctx)
}

// Close closes the connection pool.
func (r *Repository) Close() error {
	return r.db.Close()
}
