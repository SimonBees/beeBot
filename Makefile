# Makefile for beeBot Server

# Build the application
build:
	go build -o bin/server cmd/server/main.go

# Run the application
run: build
	./bin/server

# Install dependencies
deps:
	go mod tidy

# Test the application
test:
	go test ./...

# Clean build artifacts
clean:
	rm -rf bin/

# Build for different platforms
build-linux:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/server-linux cmd/server/main.go

build-windows:
	CGO_ENABLED=0 GOOS=windows go build -a -installsuffix cgo -o bin/server.exe cmd/server/main.go

build-darwin:
	CGO_ENABLED=0 GOOS=darwin go build -a -installsuffix cgo -o bin/server-darwin cmd/server/main.go

.PHONY: build run deps test clean build-linux build-windows build-darwin