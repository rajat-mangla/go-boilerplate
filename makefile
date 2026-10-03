APP_EXECUTABLE := ./out/boilerplate-service

# TESTS 		###########################################################################################

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
