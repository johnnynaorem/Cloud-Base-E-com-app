# Step 1: Build the React app using Node.js
FROM node:18-alpine AS builder

WORKDIR /app

# Copy package files and install dependencies
COPY package*.json ./
RUN npm install

# Copy the rest of the code
COPY . .

# Build the app for production
RUN npm run build

# Step 2: Serve the build using a minimal web server (e.g., nginx)
FROM nginx:alpine

# Copy build output to nginx html directory
COPY --from=builder /app/dist /usr/share/nginx/html

# Optional: replace nginx default config (for SPA support)
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Expose port 80 and start nginx
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
