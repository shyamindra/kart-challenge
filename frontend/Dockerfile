# frontend/Dockerfile

# --- Build Stage ---
FROM node:20-alpine AS builder

WORKDIR /app

# Copy package.json and package-lock.json
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the source code
COPY . .

# Build the application
# VITE_API_BASE_URL is passed via docker-compose environment variables
RUN npm run build

# --- Final Stage (Serving with Nginx) ---
FROM nginx:stable-alpine AS final

# Copy the built files from the builder stage to the Nginx public directory
COPY --from=builder /app/dist /usr/share/nginx/html

# Create a simple Nginx config for a SPA to handle routing
RUN echo "server { \
    listen 80; \
    location / { \
        root /usr/share/nginx/html; \
        index index.html; \
        try_files $uri $uri/ /index.html; \
    } \
}" > /etc/nginx/conf.d/default.conf

# Expose the port Nginx is listening on
EXPOSE 80

# Start Nginx
CMD ["nginx", "-g", "daemon off;"]
