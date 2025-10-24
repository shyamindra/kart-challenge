# Frontend Guide: React/TypeScript Application

## Tech Stack and Architecture

The frontend is a single-page application built with **React and TypeScript**, bundled using **Vite**.

*   **Framework:** React with TypeScript
*   **Build Tool:** Vite
*   **State Management:** React Context API + `useReducer` hook (in `src/context/CartContext.tsx`) for a centralized, robust shopping cart state.
*   **Styling:** Component-level CSS/CSS Modules for scoped styles.

The application follows a component-based architecture:

| Directory | Responsibility | Examples |
| :--- | :--- | :--- |
| `src/components` | Reusable UI | `ProductCard.tsx`, `CartItem.tsx`, `AddToCartButton.tsx` |
| `src/context` | Global State | `CartContext.tsx` (manages cart items, quantities, and totals) |
| `src/services` | API Interaction | `apiService.ts` (handles all communication with the backend API) |

## Containerization: Dockerfile

The `frontend/Dockerfile` uses a multi-stage build for efficient deployment:

1.  **Build Stage:** Uses `node:20-alpine` to install dependencies and run `npm run build`. The `VITE_API_BASE_URL` environment variable is set via `docker-compose.yaml` to point to the backend service.
2.  **Final Stage:** Uses a lightweight `nginx:stable-alpine` image to serve the static files from the build output. A custom Nginx configuration is included to ensure proper routing for the single-page application (SPA).

## Getting Started: Running the Web App

The frontend development server will automatically proxy API calls to the backend, so ensure the backend is running first.

1.  **Navigate to the frontend directory:**
    ```bash
    cd frontend
    ```
2.  **Install Node.js dependencies:**
    ```bash
    npm install
    ```
3.  **Start the development server:**
    ```bash
    npm run dev
    ```
    The application will typically be available at `http://localhost:5173` (or a similar port).

---

## Development Note

This codebase was developed with the assistance of **Gemini CLI** and **GitHub Copilot**. Most of the initial scaffolding, boilerplate code, and repetitive chore tasks were handled by these Large Language Models (LLMs).