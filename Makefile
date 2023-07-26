 .PHONY: all test build run clean docker help

# Binary name
BINARY=wordgame

# Versioning
VERSION="$(shell cat VERSION)"

# Go related variables.
GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin
GOFILES=$(wildcard *.go)
GOCMD=go

IMAGE="wordgame/api"
BUILDTIME=$(shell date '+%a %b %d %H:%M:%S %Y')
GITBRANCH=$(shell git rev-parse --abbrev-ref HEAD)
GITCOMMIT=$(shell git rev-parse --short HEAD)

IMAGE_TAG=${IMAGE}:${VERSION}-${GITBRANCH}-${GITCOMMIT}

INTERNAL_PKG = $(shell go list github.com/fleetdm/wordgame/internal/...)

CURRENT_DIR = $(shell pwd)

# All process
all: test build

# Build local docker image
local-docker:
	@echo Building ${IMAGE_TAG}
	@make build
	@docker build -f ${CURRENT_DIR}/docker/Dockerfile.local -t ${IMAGE_TAG} .

# Runs all tests
test:
	$(GOCMD) test -v ./...

# Run the application
run-local-dept:
	@echo "Running"
	@cd docker && docker-compose up -d

# Build the project
build:
	@echo "Building project"
	$(GOCMD) build -o $(GOBIN)/$(BINARY) -v ./cmd
	@echo "Project built"

# Run the application
run: build
	@echo "Running"
	$(GOBIN)/$(BINARY)

# Clean build cache files
clean:
	@echo "Cleaning build cache"
	@GO111MODULE=on $(GOCMD) clean
	@echo "Cleaned"

# Build docker image
docker-build:
	docker build --no-cache -t mywordgame -f ${CURRENT_DIR}/docker/Dockerfile .

# Run docker image
docker-run: docker-build
	docker run -p 1337:1337 mywordgame

# Generate swagger files
generate-swagger:
	swagger generate spec --scan-models -w ./cmd -o ./docs/swaggerui/swagger.json

# Format go code
fmt:
	@echo "Formatting project"
	@gofmt -w -e `find . -name "*.go" | grep -v vendor`
	@echo "Formatted"

# Lint go code
lint:
	@echo "Linting API" $(INTERNAL_PKG)
	@golint -set_exit_status $(INTERNAL_PKG)
	@golangci-lint run --timeout 3m0s
	@echo "Linted"

# Vet go code
vet:
	@echo "Vetting "
	@go vet ./...
