# Agent Instructions for Kart Challenge

This document outlines the specific roles and responsibilities for the AI agent (me) during the development of the Kart Challenge project.

## Backend Development
- **User Responsibility:** All backend coding and implementation will be handled by the user.
- **Agent Role:** I will **not** make any backend code changes unless explicitly specified by the user.

## Agent Responsibilities
I will provide assistance with the following tasks:
- **Chores:** General maintenance tasks, minor fixes, and utility operations.
- **Documentation:** Creating, updating, and maintaining project documentation.
- **Setup:** Assisting with project setup, environment configuration, and tool installation.
- **Committing:** Preparing and suggesting Git commits for changes made.
- **Frontend Development:** Providing direct assistance with frontend code changes, implementation, and debugging.
- **Code Review:** Reviewing code written by the user for quality, adherence to best practices, and potential issues.
- **Automated Tasks:** Setting up and running linting, tests, and other automated checks.

## Common Tasks/Shortcuts

For your convenience, here are some common tasks and the commands I would typically use to execute them. You can refer to these when asking me to perform these actions.

- **Run Lint:** `golangci-lint run ./...`
- **Run Tests:** `go test ./...`
- **Review Uncommitted Code:** `git diff HEAD`
- **Build Project:** `go build ./...`
- **Commit Changes (Interactive):**
    1.  I will first show you the current `git status`.
    2.  Then, I will show you the `git diff HEAD` (uncommitted changes).
    3.  I will ask you to confirm which files to stage (`git add ...`) and for the commit message.
    4.  Finally, I will execute `git commit -m "Your message"`."

## Code Review Guidelines

When performing a code review, I will adhere to the following principles, focusing on being friendly, collaborative, and constructive. My feedback will aim for 1 to 3 significant, actionable improvements, providing code examples where useful, and avoiding minor issues unless immediately actionable.

**General Review Principles:**
*   **Contextual Understanding:** I will analyze the PR description to understand the context and purpose of changes, looking for testing instructions and assessing alignment with the description.
*   **UI/Visual Changes:** I cannot visually verify results. My review will focus on technical implementation (e.g., CSS, HTML structure, accessibility).
*   **Functional Changes:** I will verify if changes can be tested as described. If testing instructions are missing, I will suggest what should be tested.
*   **Readability and Clarity:** Prioritize code readability. Check for clear and descriptive naming conventions for variables, functions, and components. Ensure consistent code formatting.
*   **Single Responsibility Principle (SRP):** Verify that each component or function has a single, well-defined responsibility. Avoid overly large or complex components.
*   **Avoid Duplication:** Look for and eliminate redundant code or logic.
*   **Meaningful Comments:** Ensure comments explain *why* something is done, especially for complex logic, rather than *what* is done.
*   **Avoid Magic Numbers and Strings:** Use constants or configuration files for hardcoded values.

### Technology-Specific Practices

#### Go (Backend)

**Coding Best Practices:**
*   **Simplicity and Readability:** Prioritize clear, concise, and maintainable code.
*   **Formatting:** Use `go fmt` for consistent code style.
*   **Naming Conventions:** Use descriptive names; `camelCase` for local, `PascalCase` for exported. Short, meaningful package names.
*   **Error Handling:** Treat errors as values, check immediately, return errors instead of panics, wrap errors for context.
*   **Concurrency:** Use goroutines and channels for communication; protect shared memory with `sync.Mutex`.
*   **Project Structure:** Organize by feature, use `cmd` for main packages, `internal` for private code.
*   **Interfaces:** Small, focused interfaces; depend on interfaces, return structs.
*   **Testing:** Use `go test`, separate test files.

**Logging Best Practices:**
*   **Structured Logging:** Use structured logging (e.g., JSON) with key-value pairs for machine readability and easy analysis. Prefer `log/slog` (Go 1.21+) or high-performance libraries like Zap or Zerolog.
*   **Contextual Information:** Always include relevant context (e.g., request IDs, user IDs) in logs to trace events, often passed via the `context` package.
*   **Appropriate Levels:** Use log levels (DEBUG, INFO, WARN, ERROR) correctly to control verbosity. Reserve `ERROR` for exceptional circumstances.
*   **Security:** Never log sensitive data such as passwords, API keys, or PII.
*   **Performance:** Check log levels before performing expensive operations to avoid unnecessary work. Avoid logging in tight loops or hot paths.

**Code Review Focus for Go:**
*   **PR Structure:** Review for small, focused PRs with clear descriptions.
*   **Automated Checks:** Ensure `gofmt`, `goimports`, and static analyzers (e.g., `golangci-lint`) have been run.
*   **Testing:** Verify the presence of meaningful unit/integration tests.
*   **Error Handling:** Ensure correct error handling; avoid silent discards and double reporting.
*   **Concurrency Safety:** Check goroutine management, proper synchronization, and `context` usage.
*   **Readability & Idiomacy:** Assess for clear naming, small functions, and adherence to idiomatic Go (e.g., "accept interfaces, return structs," avoid premature abstraction).
*   **Documentation:** Look for comments explaining *why* complex logic exists.

#### React & TypeScript (Frontend)

**Coding Best Practices:**
*   **Strict Typing:** Enable `"strict": true` in `tsconfig.json`.
*   **Tooling:** Integrate ESLint & Prettier for consistent formatting and quality.
*   **Type Annotations:** Explicitly type props and state using `interface` or `type`.
*   **Avoid `any`:** Minimize `any` type usage; use specific or union types.
*   **Utility Types:** Leverage TypeScript utility types (e.g., `Pick`, `Omit`).
*   **Component Structure:** Prefer functional components with hooks; keep components small and focused (Single Responsibility Principle).
*   **Naming:** Consistent and meaningful naming.

**Code Review Focus for React & TypeScript:**
*   **Readability & SRP:** Check for clear naming, consistent formatting, and components/functions adhering to a single responsibility.
*   **Type Safety:** Ensure `strict` mode is enabled, types are accurate, and `any` usage is minimized.
*   **Component Design:** Verify the use of functional components with hooks and appropriate component modularity.
*   **Performance:** Look for potential performance issues (e.g., unnecessary re-renders) and suggest optimizations.
*   **Tooling Adherence:** Confirm ESLint and Prettier have been used.

#### OpenAPI

**Coding Best Practices:**
*   **Design-First & Single Source of Truth:** Design API in OpenAPI before coding; maintain the spec as the definitive reference.
*   **Clarity & Conciseness:** Use descriptive names and clear language for paths, operations, parameters, and responses.
*   **Structure:** Specify `servers` property; define schemas globally in `components/schemas` (no inline schemas).
*   **Completeness:** Include examples for parameters, media types, and schemas; define authentication schemes; ensure every operation has a unique `operationId`; document all responses, including errors.
*   **Consistency:** Adhere to naming standards; organize endpoints with tags.
*   **Version Control:** Manage the spec with Git.
*   **Validation:** Use validation tools.

**Code Review Focus for OpenAPI:**
*   **Completeness & Accuracy:** Ensure all endpoints, parameters, responses, and security schemes are documented and accurately reflect API behavior.
*   **Consistency & Usability:** Verify naming conventions, structure, and overall ease of understanding for client generation.
*   **Validation:** Recommend using OpenAPI validators.
*   **Security:** Ensure security schemes are correctly defined.
*   **Examples:** Verify examples are present and accurate.
## Important Reminders
    - I will only review what is present in the `git diff`.
    - I will not make assumptions about code not shown in the diff.
    - I will ignore changes that appear to be generated code.
    - **Never commit `.env` files:** `.env` files contain sensitive information and should never be committed to version control. I will ensure they are added to `.gitignore`.
    - **Config file changes:** For any other configuration files, I will double-check with the user before making significant changes or committing them.

# Project Overview: Kart Challenge

This project is a mini food ordering web application, referred to as the "Kart Challenge". It features product listing and a shopping cart functionality, designed to integrate with an e-commerce API for product listing and order placement. The project includes a Go-based backend implementation and a React/TypeScript frontend built with Vite. The API definition is based on OpenAPI 3.1.

## Main Technologies

*   **Backend:** Go, `go-chi/chi` router, `oapi-codegen` for OpenAPI integration.
*   **Frontend:** React, TypeScript, Vite, ESLint, Prettier.
*   **API:** OpenAPI 3.1.

## Building and Running

### Backend

The backend is written in Go and uses a `Makefile` for common development tasks.

*   **Generate API Code:**
    ```bash
    make generate
    ```
    This command generates Go types and server interfaces from the `api/openapi.yaml` specification using `oapi-codegen`.

*   **Build Application:**
    ```bash
    make build
    ```
    Compiles the Go application.

*   **Run Tests:**
    ```bash
    make test
    ```
    Executes all Go tests.

*   **Run Linter:**
    ```bash
    make lint
    ```
    Runs `golangci-lint` for code quality checks.

*   **Clean Build Artifacts:**
    ```bash
    make clean
    ```
    Removes compiled binaries and generated files.

*   **Run the Backend Server:**
    After building, you can run the server from the `backend` directory:
    ```bash
    go run ./cmd/api/main.go
    ```
    The server will listen on port `8080` by default, or on the port specified by the `PORT` environment variable.

### Frontend

The frontend is a React application built with Vite and TypeScript.

*   **Install Dependencies:**
    ```bash
    npm install
    ```
    (Run this command in the `frontend/` directory to install all required Node.js packages.)

*   **Start Development Server:**
    ```bash
    npm run dev
    ```
    Starts the Vite development server, typically accessible at `http://localhost:5173` (or another port if 5173 is in use).

*   **Build for Production:**
    ```bash
    npm run build
    ```
    Compiles the React application for production deployment. The output will be in the `dist/` directory.

*   **Run Linter:**
    ```bash
    npm run lint
    ```
    Runs ESLint to check for code style and quality issues.

*   **Preview Production Build:**
    ```bash
    npm run preview
    ```
    Serves the production build locally for testing.

## Development Conventions

*   **API-First Approach:** The project uses an OpenAPI specification (`api/openapi.yaml`) as the central definition for the API, with code generation (`oapi-codegen`) for the Go backend.
*   **Go Modules:** Backend dependencies are managed using Go Modules.
*   **TypeScript & React:** The frontend is developed using TypeScript and the React framework, bundled with Vite.
*   **Code Quality:** ESLint and Prettier are configured for the frontend to maintain consistent code style and identify potential issues.