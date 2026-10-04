# go-boilerplate

A Go HTTP API service scaffold with configuration loading, structured logging,
middleware, PostgreSQL repository helpers, SQL migrations, and graceful shutdown.

## Requirements

- Go 1.27.1 or later
- PostgreSQL client tools (`psql`) for the database create/drop Make targets

## Getting started

Generate a sample configuration:

```sh
go run ./cmd/main generate-config --config-file sample.application.yml
```

Start the API server:

```sh
go run ./cmd/main start --config-file sample.application.yml
```

The server listens on `:8080` by default. `make run-local` starts the server
using `test.application.yml`. Configuration is loaded from YAML and then
overridden by environment variables; see `config/config.go` for available
settings.

## PostgreSQL

Database connectivity is optional for the API. Set `DB_ENABLED: true` in the
configuration (or `DB_ENABLED=true` in the environment) to open and verify the
database connection during server startup. PostgreSQL connection options are
under the `DB` YAML section and use the `DB_` environment-variable prefix.

The Makefile targets that create and drop a database use `psql` and these
variables:

| Make variable | Default |
| --- | --- |
| `DB_HOST` | `localhost` |
| `DB_PORT` | `5432` |
| `DB_USER` | `postgres` |
| `DB_NAME` | `boilerplate` |

For example, override these on the Make command line:

```sh
make db-create DB_HOST=localhost DB_USER=postgres DB_NAME=boilerplate
```

Supply the password through the usual PostgreSQL mechanisms, such as
`.pgpass` or `PGPASSWORD`. Database creation/drop Make variables must match
the corresponding connection settings in `DB_CONFIG`, which defaults to
`sample.application.yml`.

Create and apply migrations:

```sh
make db-create-migration MIGRATION_NAME=add_example_table
# Edit the generated migrations/<unix-epoch>_add_example_table.{up,down}.sql
make db-migrate DB_CONFIG=sample.application.yml
```

Available database targets:

| Target | Action |
| --- | --- |
| `make db-setup` | Create the configured database, then apply pending migrations. |
| `make db-create` | Create the database using `psql`. |
| `make db-drop` | Drop the database using `psql`. |
| `make db-migrate` | Apply pending migrations using the config in `DB_CONFIG`. |
| `make db-rollback` | Roll back the most recent migration. |
| `make db-create-migration MIGRATION_NAME=<name>` | Create paired up/down SQL files. |
| `make db-reset` | Drop and recreate the database, then apply migrations. **This deletes its data.** |

Migration execution reads the connection details from `DB_CONFIG`; it does not
require `DB_ENABLED` to be true. `db-reset` runs database operations in order
and should only be used when it is safe to discard the database contents.

## Development commands

| Target | Action |
| --- | --- |
| `make build` | Build the service into `out/boilerplate-service`. |
| `make fmt` | Format Go files with `gofumpt` and `gci` through `golangci-lint`. |
| `make lint` | Run `golangci-lint` across all packages. |
| `make test` | Run the Go tests with the race detector. |
| `make sample-config` | Generate `sample.application.yml`. |

The formatter, linter, and test runner are declared in `tools/go.mod` and
invoked through `go tool`.

## HTTP endpoints

| Method | Path | Description |
| --- | --- | --- |
| `GET` | `/ping` | Health check; returns `OK`. |
| `GET` | `/internal/v1/sample` | Example API route. |

Set `API.DEBUG_MODE: true` (or `DEBUG_MODE=true`) to enable Go profiling at
`/debug/pprof/` and the Statsviz dashboard at `/debug/statsviz/`.

## Project layout

- `cmd/main/` — CLI entry point, HTTP server, router, and migration commands
- `config/` — defaults and YAML/environment configuration loading
- `repository/` — PostgreSQL connection pool and context-aware SQL helpers
- `dbmigrate/`, `migrations/` — migration runner and SQL migration files
- `handlers/` — HTTP handlers and response helpers
- `middleware/` — request context and panic recovery
- `client/`, `service/` — client and service registries
