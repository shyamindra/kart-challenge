# Shopping Cart

Build a mini food ordering web app featuring product listing and a functional shopping cart.\
Prioritize correctness in functionality while getting it to look as close to the design as possible.

For this task you will need to integrate to our demo e-commerce API for listing products and placing orders.

**API Reference**

You can find our [API Documentation](https://orderfoodonline.deno.dev/public/openapi.html) here.

API documentation is based on [OpenAPI3.1](https://swagger.io/specification/v3/) specification.
You can also find spec file [here](https://orderfoodonline.deno.dev/public/openapi.yaml).
 
**Functional Requirements**

- Display products with images
- Add items to the cart and remove items
- Show order total correctly
- Increase or decrease item count in the cart
- Show order confirmation after placing the order
- Interactive hover and focus states for elements

**Bonus Goals**

- Allow users to enter a discount code (above the "Confirm Order" button)
- Discount code `HAPPYHOURS` applies 18% discount to the order total
- Discount code `BUYGETONE` gives the lowest priced item for free
- Responsive design based on device's screen size

**Are You a Full Stack Developer??**

Impress us by implementing your own version of the API based on the OpenAPI specification.\
Choose any language or framework of your choice. For example our top pick for backend is [Go](https://go.dev)

> The API immplementation example available to you at orderfoodonline.deno.dev/api is simplified and doesn't handle some edge cases intentionally.
> Use your best judgement to build a Robust API server.

**Checkout our [advanced backend challenge](./backend-challenge/README.md) for extra bonus points

## Design

You can find a [Figma](https://figma.com) design file `design.fig` that you can use.
You might have to use your best judgement for some mobile layout designs and spacing.

### Style Guide

The designs were created to the following widths:

- Mobile: 375px
- Desktop: 1440px

> 💡 These are just the design sizes. Ensure content is responsive and meets WCAG requirements by testing the full range of screen sizes from 320px to large screens.

**Typography**

- Font size (product names): 16px

### Font

- Family: [Red Hat Text](https://fonts.google.com/specimen/Red+Hat+Text)
- Weights: 400, 600, 700

## Getting Started

Feel free to use any tool or workflow ou are comformtable with.\
Here is an example workflow (you can use it as a reference or use your own workflow)

1. Create a new public repository on [GitHub](https://github.com) (alternatively you can use GitLab, BitBucket or Git server of your choice).
   If you are creating your repository on GitHub, you can chose to use this repository as a starting template. (Click on Use template button at the top)
2. Look through the deisngs to plan your project. This will help you design UI libraries or tools.
3. Create a [Vite](https://vite.dev) app to bootstrap a modern front-end project (alternatively use the framework of your choice).
4. Structure your HTML and preview before theming and adding interactive functionality.
5. Test and Iterate to build more features
6. Deploy your app anywhere securely. You may use AWS, Vercel, Deno Deploy, Surge, CloudFlare Pages or some other web app deployment services.
7. Additionally configure your repository to automatically publish your app on new commit push (CI).

> 💡 Replace or Modify this README to explain your solution and how to run and test it.

_By following these guidelines, you should be able to build a functional and visually appealing mini e-commerce shopping portal that meets the minimum requirements and bonus goals. Good luck! 🚀_

Feel free to use any tool or workflow ou are comformtable with.

## Getting Started: Bringing up the Application

To run the full application, you must start both the backend API server and the frontend development server.

### 1. Backend Setup (Go API)

The backend is located in the `backend/` directory.

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

### 2. Frontend Setup (React/Vite)

The frontend is located in the `frontend/` directory.

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
    The application will typically be available at `http://localhost:5173` (or a similar port) and will automatically proxy API calls to the running backend server.

### 3. Running with Docker Compose

For a fully containerized setup, you can use Docker Compose to run both the backend and frontend services simultaneously.

1.  **Ensure Docker is running** on your system.
2.  **Build and start the services** from the project root:
    ```bash
    docker compose up --build
    ```
3.  **Access the application** at `http://localhost:3000`.



## Documentation

- [Project Overview](./docs/developer_manual/SUMMARY.md): Summary of the project, its features, and the technologies used.
- [Backend Guide](./docs/developer_manual/BACKEND_GUIDE.md): Detailed guide on the Go API server's architecture and setup.
- [Frontend Guide](./docs/developer_manual/FRONTEND_GUIDE.md): Detailed guide on the React/TypeScript application's architecture and setup.

**Resources**

- API documentation: https://orderfoodonline.deno.dev/public/openapi.html
- API specification: https://orderfoodonline.deno.dev/public/openapi.yaml
- Figma design file: [design.fig](./design.fig)
- Red Hat Text font: https://fonts.google.com/specimen/Red+Hat+Text

