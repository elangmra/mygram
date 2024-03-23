# Stage 1: Build the Go application
FROM golang:1.21.0-bullseye AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o app .

# Stage 2: Create a minimal runtime image
FROM debian:buster-slim

# Set the working directory inside the container
WORKDIR /app

# Copy the executable from the builder stage
COPY --from=builder /app/app .

# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["./app"]
