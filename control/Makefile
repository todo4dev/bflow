# --- Configuration ---
ENV_FILE     ?= .env
DC_FILE      ?= compose.yml
COMPOSE      ?= docker compose
PROJECT_NAME ?= $(or $(COMPOSE_PROJECT_NAME),control)
E2E_DIR      = e2e

.PHONY: help docker-up-d docker-down e2e-install e2e-test e2e-wait test

help:
	@echo ======================================================================
	@echo BFlow Control
	@echo ======================================================================
	@echo Infrastructure (Docker):
	@echo   make docker-up-d    # Start services in background
	@echo   make docker-down    # Stop and remove containers
	@echo   make docker-build   # Rebuild images
	@echo   make docker-logs    # Follow logs
	@echo 
	@echo E2E Testing (Robot + uv):
	@echo   make e2e-install    # Sync Python dependencies
	@echo   make e2e-test       # Run tests using variables.py loader
	@echo   make e2e-clean      # Remove test results
	@echo 
	@echo Full Workflow:
	@echo   make test           # Up -> Wait -> Run E2E
	@echo ======================================================================


# --- Docker ---
docker-up-d:
	$(COMPOSE) --env-file $(ENV_FILE) -f $(DC_FILE) -p $(PROJECT_NAME) up -d

docker-down:
	$(COMPOSE) --env-file $(ENV_FILE) -f $(DC_FILE) -p $(PROJECT_NAME) down --remove-orphans

# --- E2E (uv) ---
e2e-install:
	@cd $(E2E_DIR) && uv sync

e2e-wait:
	@echo ">>> Waiting for API healthcheck..."
	@cd $(E2E_DIR) && uv run python wait_for_api.py

e2e-test: e2e-install
	@echo ">>> Running Robot Framework tests..."
	@cd $(E2E_DIR) && uv run robot -d results tests/

test: docker-up-d e2e-wait e2e-test