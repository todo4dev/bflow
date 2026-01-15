ENV_FILE ?= .env
DC_FILE  ?= compose.yml
COMPOSE  ?= docker compose

ifneq (,$(wildcard $(ENV_FILE)))
include $(ENV_FILE)
export
endif

PROJECT_NAME ?= $(or $(COMPOSE_PROJECT_NAME),control)

.DEFAULT_GOAL := help

.PHONY: help
help:
	@echo "Targets:"
	@echo "  make docker-up       # docker compose up (foreground)"
	@echo "  make docker-up-d     # docker compose up -d"
	@echo "  make docker-down     # docker compose down --remove-orphans"
	@echo "  make docker          # down + up (foreground)"
	@echo "  make docker-d        # down + up -d"
	@echo "  make docker-build    # build"
	@echo "  make docker-logs     # logs -f"
	@echo "  make docker-ps       # ps"

.PHONY: docker-up docker-up-d docker-down docker docker-d docker-build docker-logs docker-ps

docker-up:
	$(COMPOSE) --env-file $(ENV_FILE) -f $(DC_FILE) -p $(PROJECT_NAME) up

docker-up-d:
	$(COMPOSE) --env-file $(ENV_FILE) -f $(DC_FILE) -p $(PROJECT_NAME) up -d

docker-down:
	$(COMPOSE) --env-file $(ENV_FILE) -f $(DC_FILE) -p $(PROJECT_NAME) down --remove-orphans

docker: docker-down docker-up

docker-d: docker-down docker-up-d

docker-build:
	$(COMPOSE) --env-file $(ENV_FILE) -f $(DC_FILE) -p $(PROJECT_NAME) build

docker-logs:
	$(COMPOSE) --env-file $(ENV_FILE) -f $(DC_FILE) -p $(PROJECT_NAME) logs -f

docker-ps:
	$(COMPOSE) --env-file $(ENV_FILE) -f $(DC_FILE) -p $(PROJECT_NAME) ps
