BINARY    := lark-cli
MODULE    := github.com/lark-dev/lark-cli
VERSION   ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)
LDFLAGS   := -s -w -X main.version=$(VERSION)

.PHONY: build test lint install clean

build:
	go build -ldflags "$(LDFLAGS)" -o $(BINARY) ./cmd/lark-cli

test:
	go test ./... -race -count=1

lint:
	go vet ./...
	@echo "Checking formatting..."
	@test -z "$$(gofmt -l .)" || (echo "Files need gofmt:" && gofmt -l . && exit 1)

install:
	go install -ldflags "$(LDFLAGS)" ./cmd/lark-cli

clean:
	rm -f $(BINARY)
	rm -f coverage.out coverage.html
