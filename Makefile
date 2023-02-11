PROJECT_NAME:=ohlc

.PHONY: all
all: help

.PHONY: help
help:
	@echo "------------------------------------------------------------------------"
	@echo "${PROJECT_NAME}"
	@echo "------------------------------------------------------------------------"
	@grep -E '^[a-zA-Z0-9_/%\-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: serve
serve: deps-godotenv ## Run locally
	godotenv -f .env go run -race ./cmd/server

.PHONY: build
build:  ## Build application binaries
	go build -race -o ./bin/server ./cmd/server

.PHONY: deps
deps: deps-godotenv deps-moq deps-swag ## Install build dependencies

.PHONY: deps-moq
deps-moq: ## Install build dependencies: moq
	@which -s moq || go install github.com/djui/moq@v0.3.3

.PHONY: deps-godotenv
deps-godotenv: ## Install build dependencies: godotenv
	@which -s godotenv || go install github.com/joho/godotenv/cmd/godotenv@latest

.PHONY: deps-swag
deps-swag: ## Install build dependencies: Swag
	@which -s swag || go install github.com/swaggo/swag/cmd/swag@v1.8.10

.PHONY: image
image: ## Create Docker image
	docker build --no-cache -t ohlc .

.PHONY: docker-only-up
docker-only-up: ## Create Docker image
	docker run --env-file .env --publish 8090:8090 ohlc:latest

.PHONY: docker-up
docker-up: ## Start docker-compose
	docker-compose up --build --abort-on-container-exit

.PHONY: docker-down
docker-down: ## Stop docker-compose
	docker-compose down --remove-orphans

.PHONY: test
test: ## Run unit tests
	go test -v -race ./cmd/... ./internal/...
