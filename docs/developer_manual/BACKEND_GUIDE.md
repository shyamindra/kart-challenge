# Backend Guide: Go API Server

## Project Structure and Architecture

The backend is a Go application located in the `backend/` directory, built using a **layered architecture** to ensure separation of concerns, testability, and maintainability.

| Directory | Responsibility | Description |
| :--- | :--- | :--- |
| `cmd/api` | Entry Point | Contains `main.go` for application startup and dependency injection. |
| `internal/handler` | API Layer | Handles HTTP requests and responses, acting as a thin wrapper around the service layer. |
| `internal/service` | Business Logic | Contains core logic like `order_service.go` for calculating totals and applying discounts. |
| `internal/storage` | Data Access | Abstracts the database (SQLite) operations via repository interfaces. |
| `internal/promocode` | External Service | Handles the complex, pre-processed validation logic for promo codes. |

## Containerization: Dockerfile

The `backend/Dockerfile` uses a multi-stage build to create a minimal, production-ready image:

1.  **Build Stage:** Uses `golang:1.21-alpine` to compile the Go binary and execute the `scripts/prepare_data.sh` script to download and decompress the coupon base files into the `data/` directory.
2.  **Final Stage:** Uses a minimal `alpine:latest` image, copying only the compiled binary and the prepared `data/` directory.

## Getting Started: Running the API

The API server must be running before starting the frontend.

1.  **Navigate to the backend directory:**
    ```bash
    cd backend
    ```
2.  **Install dependencies and build the application:**
    ```bash
    go mod download
    go build -o api ./cmd/api
    ```
3.  **Run the API server:**
    ```bash
    ./api
    # The server should start on the port configured in main.go (e.g., :8080)
    ```
    *Note: The server uses an SQLite database for persistence. If you need to prepare initial product data, refer to the `backend/scripts/prepare_data.sh` script.*

---

## Development Note

This codebase was developed with the assistance of **Gemini CLI** and **GitHub Copilot**. Most of the initial scaffolding, boilerplate code, and repetitive chore tasks were handled by these Large Language Models (LLMs).