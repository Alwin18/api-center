FROM golang:lts-alpine AS base

FROM base AS builder

# Set the working directory
WORKDIR /app

# Copy the source code into the container
COPY . .

# Download go modules
RUN go mod download

# Tidy up the go modules
RUN go mod tidy

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main cmd/main.go

# Build a smaller image that will only contain the application's binary
FROM alpine:latest

# Create a non-root user and group
RUN addgroup -g 1001 -S api-center && adduser -u 1001 -S api-center -G api-center

# Set the working directory
WORKDIR /app

# Switch to the non-root user
USER api-center

# Copy the application's binary
COPY --from=builder /app .

# Expose the application's port
EXPOSE 9200

# Command to run the application when starting the container
CMD ["./main"]
