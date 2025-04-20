APP_NAME := echo-chamber
GO := go
PKG := ./...
PREFIX := [make]
BUILD_DIR := bin/
LOGS_DIR := logs/

# Default target
.DEFAULT_GOAL := build

# Declare phony targets
.PHONY: fmt vet build test clean check run

fmt:
	@echo "$(PREFIX) Formatting source code..."
	@$(GO) fmt $(PKG)

vet: fmt
	@echo "$(PREFIX) Running vet to check code..."
	@$(GO) vet $(PKG)

build: vet
	@echo "$(PREFIX) Building $(APP_NAME) in $(BUILD_DIR)..."
	@mkdir -p $(BUILD_DIR)
	@$(GO) build -o $(BUILD_DIR)$(APP_NAME) ./cmd/server/main.go

test:
	@echo "$(PREFIX) Running tests..."
	@$(GO) test $(PKG) -v

clean:
	@echo "$(PREFIX) Cleaning build artifacts and logs..."
	@rm -rfv $(BUILD_DIR)
	@rm -rfv $(LOGS_DIR)

check: fmt vet test
	@echo "$(PREFIX) Quality checks (format, vet, tests) complete!"

run: build
	@echo "$(PREFIX) Running $(APP_NAME)..."
	@$(BUILD_DIR)$(APP_NAME)