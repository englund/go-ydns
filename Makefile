include .env

# Binary output name
BINARY = ydns-updater

# Go related variables
GOBASE = $(shell pwd)
GOBIN = $(GOBASE)/bin

# Go commands
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test

all: build

build: 
	$(GOBUILD) -o $(GOBIN)/$(BINARY) -v

build-arm64:
	GOOS=linux GOARCH=arm64 $(GOBUILD) -o $(GOBIN)/$(BINARY) -v

test: 
	$(GOTEST) -v ./...

clean: 
	$(GOCLEAN)
	rm -f $(GOBIN)/$(BINARY)

deploy: build-arm64
	scp $(GOBIN)/$(BINARY) $(SSH_USER)@$(SSH_HOST):$(SSH_PATH)/$(BINARY)

.PHONY: all build build-arm64 test clean deploy