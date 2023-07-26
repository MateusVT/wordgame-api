# Dockerfile
# Step 1: Build the application.
FROM golang:1.20-alpine AS builder

# Install git and make
# Git is required for fetching the dependencies.
# Make is required for building the application using the Makefile.
RUN apk update && apk add --no-cache git make

WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Copy the configuration file from the local filesystem to the filesystem of the container
COPY ./config.local.yml /app/config.local.yml

# Build the application
RUN make build


# Step 2: Create lean production image.
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/bin/wordgame .

# Expose port 8080 to the outside
EXPOSE 8080

# Command to run the executable
CMD ["./wordgame"]
