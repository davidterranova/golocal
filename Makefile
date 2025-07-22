BINARY=golocal

# system variables
TARGET_DIR=target
BRANCH ?= $(shell git branch | grep "^\*" | sed 's/^..//')
COMMIT ?= $(shell git rev-parse --short HEAD)
VERSION=0.1-$(COMMIT)
BUILDTIME ?= $(shell date -u +%FT%T)

LINT ?= golangci-lint

# build flags
GOOS ?= darwin
GOARCH ?= arm64
BUILD_ENV=CGO_ENABLED=0
LDFLAGS=-ldflags='-w -s -X github.com/davidterranova/golocal/cmd.Version=${VERSION} -X github.com/davidterranova/golocal/cmd.BuildTime=${BUILDTIME}'
BUILD_FLAGS=-a
GOBUILD=go build

# build
.PHONY: build-docker
build-docker:
	docker build -t golocal:$(VERSION) .
	docker tag golocal:$(VERSION) golocal:latest
	docker tag golocal:latest terranovadavid/golocal:latest
	docker push terranovadavid/golocal:latest

.PHONY: build
build: clean prepare
	$(BUILD_ENV) $(GOBUILD) $(BUILD_FLAGS) $(LDFLAGS) -o $(TARGET_DIR)/$(BINARY)

.PHONY: clean
clean:
	rm -rf $(TARGET_DIR)

.PHONY: prepare
prepare:
	mkdir -p $(TARGET_DIR)


# development
.PHONY: lint
lint:
	$(LINT) run ./...

.PHONY: lint-fix
lint-fix:
	$(LINT) run --fix ./...

.PHONY: test-unit
test-unit:
	find . -name '.sequence' -type d | xargs rm -rf
	go test ./... -v -count=1 -race -cover

# it requires having a mockgen installed. See: https://github.com/golang/mock
.PHONY: mockgen
mockgen:
	go generate ./...