.PHONY: all build install run clean

# Binary name
BINARY_NAME=task-tracker
BUILD_DIR=bin

all: build

# Ensure build directory exists
$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

# Build the binary
build: $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/cli-task-tracker

# Install the binary to GOPATH/bin
install:
	go install ./cmd/cli-task-tracker

# Run without installing
run:
	go run ./cmd/cli-task-tracker

# Clean build files
clean:
	rm -rf $(BUILD_DIR)
	go clean 