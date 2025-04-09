# file: Makefile
# description: Makefile for building and testing the go-input package

# Project variables
PROJECT_NAME := goinput
VERSION := 1.0.0
MODULE := github.com/AndrewDonelson/go-input

# Go commands
GO := go
GOFMT := gofmt
GOTEST := $(GO) test
GOBUILD := $(GO) build

# Build flags
LDFLAGS := -ldflags="-s -w -X main.version=$(VERSION)"
COVERAGE_FLAGS := -coverprofile=coverage.out -covermode=atomic

# Platforms to build for
PLATFORMS := windows linux darwin
OS_ARCHS := amd64 arm64

# Directories
BIN_DIR := bin
DIST_DIR := dist
CMD_DIR := cmd/goinput

# Default target
all: clean setup format test build

help:
	@echo "go-input $(VERSION) - Build Targets:"
	@echo "all        - Run clean, setup, format, test, build"
	@echo "setup      - Set up build directories and dependencies"
	@echo "clean      - Remove build artifacts and temporary files"
	@echo "format     - Format Go code with gofmt"
	@echo "test       - Run all tests"
	@echo "coverage   - Generate test coverage report"
	@echo "build      - Build for current platform"
	@echo "build-all  - Build for all supported platforms"

setup:
	@echo "Setting up project..."
	@mkdir -p $(BIN_DIR) $(DIST_DIR)
	$(GO) mod tidy

clean:
	@echo "Cleaning project..."
	@rm -rf $(BIN_DIR)
	@rm -rf $(DIST_DIR)
	@rm -f coverage.out coverage.html

format:
	@echo "Formatting code..."
	$(GOFMT) -w -s .

test:
	@echo "Running tests..."
	$(GOTEST) -v ./...

coverage:
	@echo "Generating coverage report..."
	$(GOTEST) $(COVERAGE_FLAGS) ./...
	$(GO) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated at coverage.html"

build:
	@echo "Building for current platform..."
	@mkdir -p $(BIN_DIR)
	$(GOBUILD) $(LDFLAGS) -o $(BIN_DIR)/$(PROJECT_NAME) ./$(CMD_DIR)
	@echo "Build complete: $(BIN_DIR)/$(PROJECT_NAME)"

build-all:
	@echo "Building for all platforms..."
	@mkdir -p $(DIST_DIR)
	@for platform in $(PLATFORMS); do \
		for arch in $(OS_ARCHS); do \
			output_name=$(DIST_DIR)/$(PROJECT_NAME)-$${platform}-$${arch}$(if $(findstring windows,$${platform}),.exe,); \
			echo "Building for $${platform}/$${arch} -> $${output_name}"; \
			GOOS=$${platform} GOARCH=$${arch} $(GOBUILD) $(LDFLAGS) -o $${output_name} ./$(CMD_DIR); \
		done; \
	done
	@echo "All builds complete in $(DIST_DIR)/"

.PHONY: all setup clean format test coverage build build-all help