# --- Variáveis de Configuração ---
ENV_FILE     ?= .env
DC_FILE      ?= compose.yml
COMPOSE      ?= docker compose
PROJECT_NAME ?= $(or $(COMPOSE_PROJECT_NAME),control)
E2E_DIR      = control/e2e
API_URL      ?= http://localhost:8080/health # Ajuste para seu endpoint de healthcheck

# Carrega o .env se existir
ifneq (,$(wildcard $(ENV_FILE)))
include $(ENV_FILE)
export
endif

.DEFAULT_GOAL := help

# --- Targets Auxiliares ---
.PHONY: help
help:
	@echo "Comandos Docker:"
	@echo "  make docker-up         # Sobe containers (foreground)"
	@echo "  make docker-up-d       # Sobe containers (background)"
	@echo "  make docker-down       # Para e remove containers"
	@echo "  make docker-build      # Build das imagens"
	@echo "  make docker-logs       # Tail dos logs"
	@echo ""
	@echo "Comandos E2E (Robot + uv):"
	@echo "  make e2e-install       # Instala/Sincroniza dependências com uv"
	@echo "  make e2e-test          # Executa os testes E2E"
	@echo "  make e2e-clean         # Limpa relatórios de testes"
	@echo ""
	@echo "Fluxos Completos:"
	@echo "  make test              # Sobe ambiente, espera API e roda testes"

# --- Targets Docker ---
.PHONY: docker-up docker-up-d docker-down docker-build docker-logs docker-ps

docker-up:
	$(COMPOSE) --env-file $(ENV_FILE) -f $(DC_FILE) -p $(PROJECT_NAME) up

docker-up-d:
	$(COMPOSE) --env-file $(ENV_FILE) -f $(DC_FILE) -p $(PROJECT_NAME) up -d

docker-down:
	$(COMPOSE) --env-file $(ENV_FILE) -f $(DC_FILE) -p $(PROJECT_NAME) down --remove-orphans

docker-build:
	$(COMPOSE) --env-file $(ENV_FILE) -f $(DC_FILE) -p $(PROJECT_NAME) build

docker-logs:
	$(COMPOSE) --env-file $(ENV_FILE) -f $(DC_FILE) -p $(PROJECT_NAME) logs -f

docker-ps:
	$(COMPOSE) --env-file $(ENV_FILE) -f $(DC_FILE) -p $(PROJECT_NAME) ps

# --- Targets E2E (uv) ---
.PHONY: e2e-install e2e-test e2e-clean e2e-wait

e2e-install:
	@echo ">>> Sincronizando dependências Python com uv..."
	@cd $(E2E_DIR) && uv sync

e2e-wait:
	@echo ">>> Aguardando API em $(API_URL) ficar disponível..."
	@for i in {1..20}; do \
		curl -s $(API_URL) > /dev/null && break || \
		(echo "Tentativa $$i/20: API ainda não respondeu..." && sleep 2); \
	done

e2e-test: e2e-install
	@echo ">>> Executando testes Robot Framework..."
	@cd $(E2E_DIR) && uv run robot -d results tests/

e2e-clean:
	@rm -rf $(E2E_DIR)/results

# --- Fluxo de Teste Unificado ---
.PHONY: test
test:
	@$(MAKE) docker-up-d
	@$(MAKE) e2e-wait
	@$(MAKE) e2e-test
	@# $(MAKE) docker-down # Descomente se quiser derrubar tudo após o teste