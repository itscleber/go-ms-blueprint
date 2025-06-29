ifneq (,$(wildcard ./.env))
	include .env
	export
endif

up:
	docker-compose up --build --force-recreate -d

run:
	docker-compose exec app go run ./cmd/main.go

test:
	docker-compose exec app go test ./tests

down:
	docker-compose down

pre-commit-install:
	pre-commit install

pre-commit:
	pre-commit run --all-files

lint:
	golangci-lint run ./...

test:
	docker-compose exec app go test -v ./...

cover:
	docker-compose exec app go test -coverprofile=coverage.out ./...
	docker-compose exec app go tool cover -html=coverage.out
