# go-boilerplate

A Go HTTP API service scaffold with configuration loading, structured logging,
middleware, sample routes, and graceful shutdown.

## Requirements

- Go 1.27.1 or later

## Getting started

Generate a default configuration file:

```sh
go run ./cmd/main generate-config -c application.yml
```

Start the API server with that configuration:

```sh
go run ./cmd/main start -c application.yml
```

The server listens on `:8080` by default. You can also use the included
`sample.application.yml`:

```sh
go run ./cmd/main start -c sample.application.yml
```

Configuration can be loaded from YAML and overridden with environment
variables. See `config/config.go` for the supported settings and environment
variable names.

## PostgreSQL

PostgreSQL is optional for the API server. 
Set `DB_ENABLED` to `true` in the YAML configuration or environment to open and verify a connection during startup. 

Migration commands use the configured PostgreSQL connection and the
`migrations/` directory. Run them from the repository root:

```sh
go run ./cmd/main migrate-create add_example_table
go run ./cmd/main migrate-up -c application.yml
go run ./cmd/main migrate-down -c application.yml
```

`migrate-down` rolls back one migration. Edit the generated `.up.sql` and
`.down.sql` files before applying the migration. `DB_ENABLED` controls API
startup only; explicit migration commands work regardless of its value.

## HTTP endpoints

| Method | Path | Description |
| --- | --- | --- |
| `GET` | `/ping` | Health check; returns `OK`. |
| `GET` | `/internal/v1/sample` | Example API route. |

Set `API.DEBUG_MODE` to `true` in the YAML config, or set `DEBUG_MODE=true`,
to enable the Go profiling endpoints under `/debug/pprof/` and the Statsviz
dashboard at `/debug/statsviz/`.

## Project layout

- `cmd/main/` — CLI entry point, router setup, and HTTP server
- `config/` — application configuration, defaults, and YAML/environment loading
- `repository/` — PostgreSQL connection pool and context-aware SQL helpers
- `dbmigrate/`, `migrations/` — migration runner and SQL migration files
- `handlers/` — HTTP handlers and response helpers
- `middleware/` — request context and panic recovery middleware
- `client/`, `service/` — client and service registries
