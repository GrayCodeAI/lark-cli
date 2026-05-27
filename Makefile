BINARY=lark-cli

.PHONY: build test install clean

build:
	go build -o $(BINARY) ./cmd/lark-cli

test:
	go test ./... -race -count=1

install:
	go install ./cmd/lark-cli

clean:
	rm -f $(BINARY)
