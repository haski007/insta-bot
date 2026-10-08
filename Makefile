# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
CMD_PATH=./cmd/app/
BINARY_PATH=./build/
BINARY_NAME=insta-bot
BINARY_UNIX=$(BINARY_NAME)_unix

all: test build

test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	@$(GORUN) $(CMD_PATH)*.go
run-local:
	@$(GORUN) $(CMD_PATH)*.go -c config/local.yaml -metrics_server_addr :9091
deps:
	$(GOGET) -u ./...


# Cross compilation
.PHONY: build
build:
	CGO_ENABLED=0 \
		$(GOBUILD) \
			-installsuffix cgo \
			-o $(BINARY_PATH)$(BINARY_NAME) \
			-ldflags "-X main.Version=$(APP_VERSION)" \
			$(CMD_PATH)*.go

build-linux:
	CGO_ENABLED=0 \
	GOOS=linux \
	GOARCH=amd64 \
		$(GOBUILD) \
			-installsuffix cgo \
			-o $(BINARY_PATH)$(BINARY_NAME).linux \
			-ldflags "-X main.Version=$(APP_VERSION)" \
			$(CMD_PATH)*.go

# Docker Compose: rebuild & restart with latest changes
COMPOSE ?= docker compose

.PHONY: restart restart-bot restart-instloader restart-grafana restart-prometheus restart-loki restart-promtail restart-redis
restart: ## rebuild & restart all services
	$(COMPOSE) up -d --build --force-recreate

restart-bot: ## rebuild & restart bot only
	$(COMPOSE) up -d --build --force-recreate bot

restart-instloader: ## rebuild & restart instloader only
	$(COMPOSE) up -d --build --force-recreate instloader

restart-grafana: ## recreate grafana (dashboards/provisioning)
	$(COMPOSE) up -d --force-recreate grafana

restart-prometheus: ## recreate prometheus (config)
	$(COMPOSE) up -d --force-recreate prometheus

restart-loki: ## recreate loki (config)
	$(COMPOSE) up -d --force-recreate loki

restart-promtail: ## recreate promtail (config)
	$(COMPOSE) up -d --force-recreate promtail

restart-redis: ## recreate redis
	$(COMPOSE) up -d --force-recreate redis

