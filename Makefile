# Go variables
GO=go
PKG=main
PROJECT_PATH := $(PWD)

# Go build flags
BUILD_FLAGS=-v

# Targets
all: builder

linkBinary:
	sudo ln -s $(PROJECT_PATH)/builder /usr/bin/bdev

builder:
	$(GO) build $(BUILD_FLAGS) -o builder builder.go client.go logger.go interpreter.go registry.go systemd.go environment.go

demos: registryDemo interpreterDemo builderUiDemo

tests:
	go test -v systemd_test.go systemd.go

registryDemo:
	$(GO) build $(BUILD_FLAGS) -o registryDemo registryDemo.go registry.go
	mkdir -p demo
	mv registryDemo ./demo

interpreterDemo:
	$(GO) build $(BUILD_FLAGS) -o interpreterDemo interpreterDemo.go client.go logger.go interpreter.go registry.go
	mkdir -p demo
	mv interpreterDemo ./demo

builderUiDemo:
	$(GO) build $(BUILD_FLAGS) -o builderUiDemo builderUi.go
	mkdir -p demo
	mv builderUiDemo ./demo

clean:
	rm builder
	rm update
	rm demo/registryDemo
	rm demo/builderUiDemo
	rm demo/interpreterDemo

.PHONY: builder clean registryDemo interpreterDemo builderUiDemo tests
