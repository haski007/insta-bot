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





