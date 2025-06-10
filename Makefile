.PHONY: all build install run clean uninstall

# binary names
BINARY_NAME=btl
BINARY_PATH=/usr/local/bin
LOCAL_BIN=bin/$(BINARY_NAME)

all: build

bin:
	mkdir -p bin

# build binary
build: bin
	go build -o $(LOCAL_BIN) ./cmd/cli-task-tracker
	@echo "Binary built as: $(LOCAL_BIN)"
	@echo "Run locally with: ./$(LOCAL_BIN)"
	@echo "Or install system-wide with: make install"

# install the binary to /usr/local/bin (requires sudo)
install: build
	@echo "Installing $(BINARY_NAME) to $(BINARY_PATH)..."
	@if [ -w "$(BINARY_PATH)" ]; then \
		cp $(LOCAL_BIN) $(BINARY_PATH)/; \
	else \
		echo "Need sudo permission to install to $(BINARY_PATH)"; \
		sudo cp $(LOCAL_BIN) $(BINARY_PATH)/; \
	fi
	@echo "Installation complete. Run '$(BINARY_NAME)' to get started."

# run w/o installing
run:
	go run ./cmd/cli-task-tracker

clean:
	rm -rf bin
	go clean

# uninstall binary
uninstall:
	@echo "Removing $(BINARY_NAME) from $(BINARY_PATH)..."
	@if [ -w "$(BINARY_PATH)" ]; then \
		rm -f $(BINARY_PATH)/$(BINARY_NAME); \
	else \
		echo "Need sudo permission to remove from $(BINARY_PATH)"; \
		sudo rm -f $(BINARY_PATH)/$(BINARY_NAME); \
	fi
	@echo "Uninstallation complete." 