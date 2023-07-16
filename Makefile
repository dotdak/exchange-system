
all: help

help:
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  ${YELLOW}%-30s${GREEN}%s${RESET}\n", $$1, $$2}' $(MAKEFILE_LIST)


generate: ## Generate protobuf
	buf generate
	wire ./...

lint: ## Run lint
	buf lint

wire:
	wire ./...

BUF_VERSION:=1.23.1

install: ## Install prerequisites software for development
	go install \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
		github.com/envoyproxy/protoc-gen-validate \
		github.com/google/wire/cmd/wire
	curl -sSL \
    	"https://github.com/bufbuild/buf/releases/download/v${BUF_VERSION}/buf-$(shell uname -s)-$(shell uname -m)" \
    	-o "$(shell go env GOPATH)/bin/buf" && \
  	chmod +x "$(shell go env GOPATH)/bin/buf"

build:  ## Build binary
	@./scripts/build.sh

unit-test: ## Run unit tests
	go test -v -race ./... $(OUTPUT_OPTIONS)

integration-test-start: ## Run integration test
	go test -tags=integration ./... -v -count=1

integration-test-setup: ## Setup integration test infra
	docker-compose -f docker-compose.it.yml up -d

integration-test-teardown: ## Teardown integration test infra
	docker-compose -f docker-compose.it.yml down -v	
