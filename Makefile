# Go variables
GO=go
PKG=main
PROJECT_PATH := $(PWD)

# Go build flags
BUILD_FLAGS=-v

# Targets
all: builder

linkBinary:
	sudo ln -s $(PROJECT_PATH)/bin/builder /usr/bin/bdev

builder:
	$(GO) build -o $(PROJECT_PATH)/bin/builder $(BUILD_FLAGS) $(PROJECT_PATH)/cmd/builder

tests:
	$(GO) test ./internal/lexer $(BUILD_FLAGS)
	$(GO) test ./internal/parser $(BUILD_FLAGS)
	$(GO) test ./internal/interpreterV2 $(BUILD_FLAGS)

clean:
	rm bin/builder
.PHONY: builder clean tests
