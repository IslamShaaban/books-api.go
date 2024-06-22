# Makefile

# Project-specific variables
PROJECT_NAME = books-api

# Docker Compose commands
DOCKER_COMPOSE = docker compose
DC_BUILD = $(DOCKER_COMPOSE) build
DC_UP = $(DOCKER_COMPOSE) up --build -d
DC_DOWN = $(DOCKER_COMPOSE) down
DC_RESTART = $(DOCKER_COMPOSE) restart
DC_LOGS = $(DOCKER_COMPOSE) logs
DC_LOGS_FOLLOW = $(DOCKER_COMPOSE) logs -f

# Default target
.PHONY: default
default: help

# Show help
.PHONY: help
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  up         Start and build containers"
	@echo "  down       Stop and remove containers"
	@echo "  restart    Restart the containers"
	@echo "  logs       Show logs"
	@echo "  logs-f     Follow logs"
	@echo ""

# Start and build containers
.PHONY: build
build:
	$(DC_BUILD)

# Start and build containers
.PHONY: up
up:
	$(DC_UP)

# Stop and remove containers
.PHONY: down
down:
	$(DC_DOWN)

# Restart the containers
.PHONY: restart
restart:
	$(DC_RESTART)

# Show logs
.PHONY: logs
logs:
	$(DC_LOGS)

# Follow logs
.PHONY: logs-f
logs-f:
	$(DC_LOGS_FOLLOW)
