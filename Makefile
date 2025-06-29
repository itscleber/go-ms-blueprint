ifneq (,$(wildcard ./.env))
	include .env
	export
endif

COMPOSE_FILE=docker-compose.dev.yaml

up:
	docker-compose -f $(COMPOSE_FILE) up --build --force-recreate -d

run:
	docker-compose -f $(COMPOSE_FILE) exec app go run ./cmd/main.go

test:
	docker-compose -f $(COMPOSE_FILE) exec app go test ./tests

down:
	docker-compose -f $(COMPOSE_FILE) down

pre-commit-install:
	pre-commit install

pre-commit:
	pre-commit run --all-files

lint:
	golangci-lint run ./...

test:
	go test -v ./...

cover:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
