ifneq (,$(wildcard ./.env))
	include .env
	export
endif

ifeq ($(ENV),dev)
	COMPOSE_FILE=docker-compose.dev.yaml
else ifeq ($(ENV),prod)
	COMPOSE_FILE=docker-compose.prod.yaml
else
	$(error Environment variable ENV not set or not valid)
endif

up:
	docker-compose -f $(COMPOSE_FILE) up --build --force-recreate -d

run:
	docker-compose -f $(COMPOSE_FILE) exec app go run ./cmd/main.go

test:
	docker-compose -f $(COMPOSE_FILE) exec app go test ./tests

down:
	docker-compose -f $(COMPOSE_FILE) down
