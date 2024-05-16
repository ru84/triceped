# Variables
APP_NAME = triceped
BUILD_DIR = ./bin
GO_FILES = $(wildcard *.go)
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

# Targets
.PHONY: all build run clean test fmt lint

all: build ## Default target: Build the application

build: $(GO_FILES) ## Build the application
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BUILD_DIR)/$(APP_NAME) ./cmd/...

run: build ## Run the application
	@echo "Running $(APP_NAME)..."
	@$(BUILD_DIR)/$(APP_NAME)

clean: ## Clean build artifacts
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)

test: ## Run tests
	@echo "Running tests..."
	@go test ./...

fmt: ## Format the code
	@echo "Formatting code..."
	@gofmt -s -w $(GO_FILES)

lint: ## Lint the code
	@echo "Linting code..."
	@golangci-lint run
