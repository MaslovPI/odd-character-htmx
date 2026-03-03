# Variables
BINARY_NAME=odd-character-htmx
DOCKER_IMAGE=maslovpi/odd-character-htmx:latest

.PHONY: build-all
## build-all: Compiles the Go binary for Linux and builds the Docker image
build-all:
	@echo "Building Go binary for Linux/AMD64..."
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME) .
	
	@echo "Building Docker image..."
	docker build -t $(DOCKER_IMAGE) .
	
	@echo "Cleaning up binary..."
	rm $(BINARY_NAME)

.PHONY: clean
clean:
	rm -f $(BINARY_NAME)
