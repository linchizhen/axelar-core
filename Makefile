PACKAGES=$(shell go list ./... | grep -v '/simulation')

VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=NewApp \
	-X github.com/cosmos/cosmos-sdk/version.ServerName=scavengeD \
	-X github.com/cosmos/cosmos-sdk/version.ClientName=scavengeCLI \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) 

BUILD_FLAGS := -ldflags '$(ldflags)'

.PHONY: all
all: install

.PHONY: install
install: go.sum
		go install -mod=readonly $(BUILD_FLAGS) ./cmd/scavengeD
		go install -mod=readonly $(BUILD_FLAGS) ./cmd/scavengeCLI

go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		GO111MODULE=on go mod verify

# Uncomment when you have some tests
# test:
# 	@go test -mod=readonly $(PACKAGES)
.PHONY: lint
# look into .golangci.yml for enabling / disabling linters
lint:
	@echo "--> Running linter"
	@golangci-lint run
	@go mod verify

.PHONY: build
build: go.sum
		go build -o ./build/scavengeD -mod=readonly $(BUILD_FLAGS) ./cmd/scavengeD
		go build -o ./build/scavengeCLI -mod=readonly $(BUILD_FLAGS) ./cmd/scavengeCLI

.PHONY: docker
docker:
	docker-compose -f docker-compose-build.yml up