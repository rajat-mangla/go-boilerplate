.PHONY: all fmt lint generate-mocks delete-mocks
all: build fmt lint test

# Application variables
APP_EXECUTABLE= ./out/boilerplate-service
DB_CONFIG ?= sample.application.yml
DB_HOST ?= localhost
DB_PORT ?= 5432
DB_USER ?= postgres
DB_NAME ?= boilerplate
MIGRATION_NAME ?=

# Tools variables
GO_TOOL = go tool --modfile=tools/go.mod
TOOLS_DIR = $(abspath ./.tools)

# DATABASE	############################################################################################
db-setup: db-create db-migrate ##@database applies all pending database migrations
db-reset: db-drop db-create db-migrate ##@database rebuilds the configured database and applies migrations

db-create: ##@database creates the configured PostgreSQL database
	psql -h "$(DB_HOST)" -p "$(DB_PORT)" -U "$(DB_USER)" -d postgres -v ON_ERROR_STOP=1 -v db_name="$(DB_NAME)" -c 'CREATE DATABASE :"db_name"'

db-drop: ##@database drops the configured PostgreSQL database
	psql -h "$(DB_HOST)" -p "$(DB_PORT)" -U "$(DB_USER)" -d postgres -v ON_ERROR_STOP=1 -v db_name="$(DB_NAME)" -c 'DROP DATABASE IF EXISTS :"db_name"'

db-migrate: ##@database applies all pending database migrations
	go run ./cmd/main migrate:run --config-file $(DB_CONFIG)

db-rollback: ##@database rolls back the most recent database migration
	go run ./cmd/main migrate:rollback --config-file $(DB_CONFIG)

db-create-migration: ##@database creates an up/down migration pair (set MIGRATION_NAME)
	@test -n "$(MIGRATION_NAME)" || (echo "Usage: make db-create-migration MIGRATION_NAME=<name>"; exit 1)
	go run ./cmd/main migrate:create "$(MIGRATION_NAME)"

# TESTS 		###########################################################################################

fmt: ##@development formats code using golangci-lint fmt (gofumpt + gci)
	$(GO_TOOL) golangci-lint fmt --enable gofumpt --enable gci

lint: ##@quality runs golangci-lint
	$(GO_TOOL) golangci-lint run ./...

test: ##@tests runs tests
	go mod tidy
	$(GO_TOOL) gotest -race -v ./...

# DEVELOPMENT	###########################################################################################
build:
	mkdir -p out
	go build -o $(APP_EXECUTABLE) cmd/main/*.go

run-local: ##@development runs a local server directly from source
	go run cmd/main/*.go start --config-file test.application.yml

sample-config: ##@development generates sample.application.yml
	go run cmd/main/*.go generate-config --config-file sample.application.yml
