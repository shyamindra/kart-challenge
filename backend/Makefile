GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOLINT=$(shell go env GOPATH)/bin/golangci-lint
GOGEN=$(shell $(GOCMD) env GOPATH)/bin/oapi-codegen # Use shell to get GOPATH for oapi-codegen

PROJECT_DIR=$(shell pwd)

.PHONY: all generate build test lint clean

all: build

generate:
	@echo "Generating Go types and server interface..."
	$(GOGEN) --config=$(PROJECT_DIR)/types.codegen.yaml ../api/openapi.yaml
	$(GOGEN) --config=$(PROJECT_DIR)/server.codegen.yaml ../api/openapi.yaml

build:
	@echo "Building project..."
	$(GOBUILD) -v ./...

test:
	@echo "Running tests..."
	$(GOTEST) -v ./...

lint:
	@echo "Running golangci-lint..."
	$(GOLINT) run ./...

clean:
	@echo "Cleaning up..."
	$(GOCMD) clean
	rm -f $(PROJECT_DIR)/internal/handler/*.gen.go
