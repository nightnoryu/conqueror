export CGO_ENABLED=0

all: build

.PHONY: modules
modules:
	go mod tidy

.PHONY: build
build: modules
	GOOS=linux GOARCH=amd64 go build -o bin/ ./cmd/conqueror

.PHONY: clean
clean:
	rm -rf bin

.PHONY: test
test:
	go test ./...

.PHONY: check
check:
	golangci-lint run
