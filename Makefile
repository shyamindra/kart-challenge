# ==============================================================================
# Go Backend Targets
# ==============================================================================

GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOLINT=$(shell go env GOPATH)/bin/golangci-lint
GOGEN=$(shell $(GOCMD) env GOPATH)/bin/oapi-codegen

.PHONY: all backend-generate backend-build backend-test backend-lint backend-clean

all: backend-build

backend-generate:
	$(GOGEN) --config=backend/types.codegen.yaml api/openapi.yaml
	$(GOGEN) --config=backend/server.codegen.yaml api/openapi.yaml

backend-build:
	$(GOBUILD) -o backend/api ./backend/cmd/api

backend-test:
	cd backend && $(GOTEST) -v ./...

backend-lint:
	cd backend && $(GOLINT) run ./...

backend-clean:
	$(GOCMD) clean ./backend
	rm -f backend/api
	rm -f backend/internal/handler/*.gen.go

# ==============================================================================
# Frontend Targets
# ==============================================================================

.PHONY: frontend-install frontend-dev frontend-build frontend-lint frontend-preview frontend-clean

frontend-install:
	npm install --prefix frontend

frontend-dev:
	npm run dev --prefix frontend

frontend-build:
	npm run build --prefix frontend

frontend-lint:
	npm run lint --prefix frontend

frontend-preview:
	npm run preview --prefix frontend

frontend-clean:
	rm -rf frontend/node_modules frontend/dist

# ==============================================================================
# Combined Targets
# ==============================================================================

.PHONY: setup dev run clean

setup: frontend-install backend-generate

dev: backend-generate
	@echo "Backend code generated. To start the full development stack, run these two commands in separate terminals:"
	@echo "1. Backend (for hot-reloading): go run ./backend/cmd/api/main.go"
	@echo "2. Frontend: npm run dev --prefix frontend"

run: backend-build
	@echo "Starting backend server in background (PID: $$!)..."
	./backend/api &
	npm run dev --prefix frontend

clean: backend-clean frontend-clean