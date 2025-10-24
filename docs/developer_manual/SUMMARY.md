# Developer Manual: Project Overview

## Project Summary

The **Shopping Cart Challenge** is a full-stack project designed to build a mini food ordering web application. It features a product listing page and a fully functional shopping cart, integrating with a custom-built e-commerce API.

The project is split into two main components:

1.  **Backend (`backend/`):** An API server implemented in **Go**, following a layered architecture (Handler, Service, Storage) and conforming to the provided OpenAPI specification. It includes custom logic for promo code validation and order processing.
2.  **Frontend (`frontend/`):** A single-page application built with **React and TypeScript** using **Vite**. It utilizes the Context API for state management and is designed to be fully responsive.

## Key Features

*   **Product Management:** Display products with images, names, and prices.
*   **Shopping Cart:** Add, remove, and adjust the quantity of items in the cart.
*   **Order Processing:** Calculate the order subtotal and final total.
*   **Discount Logic:** Support for two specific promo codes:
    *   `HAPPYHOURS`: Applies an 18% discount.
    *   `BUYGETONE`: Gives the lowest-priced item for free.
*   **User Experience:** Show an order confirmation after a successful order placement.
*   **Design:** Responsive layout targeting mobile (375px) and desktop (1440px) widths, adhering to the provided Figma design.

## Getting Started: Bringing up the Application

The recommended way to run the full application stack is using **Docker Compose**. This handles the build, networking, and data preparation for both services automatically.

### Running with Docker Compose

1.  **Ensure Docker is running** on your system.
2.  **Build and start the services** from the project root:
    ```bash
    docker compose up --build
    ```
3.  **Access the application** at `http://localhost:3000`.

For manual setup instructions, see the dedicated guides below.

## Development Note

This codebase was developed with the assistance of **Gemini CLI** and **GitHub Copilot**. Most of the initial scaffolding, boilerplate code, and repetitive chore tasks were handled by these Large Language Models (LLMs).