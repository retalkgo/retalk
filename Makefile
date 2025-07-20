VERSION := $(shell git describe --tags --abbrev=0 --match 'v*' 2>/dev/null || echo "v0.0.0")
COMMIT := $(shell git rev-parse --short HEAD)
EXTERNAL_VERSION :=

all: app

install-dev:
	go install github.com/cosmtrek/air@latest
	go install github.com/swaggo/swag/cmd/swag@latest
	go install mvdan.cc/gofumpt@latest
	go install golang.org/x/tools/cmd/goimports@latest

app:
	go build -o bin/retalk \
		-ldflags "-w -s -X 'github.com/retalkgo/retalk/internal/version.Version=$(VERSION)$(EXTERNAL_VERSION)' \
		-X 'github.com/retalkgo/retalk/internal/version.Commit=$(COMMIT)'" -v

dev:
	air

fmt:
	gofumpt -l -w .
	swag fmt
	goimports -w .

apidoc:
	swag init -g server/server.go -o internal/docs --parseDependency --parseInternal

test:
	go test ./...

.PHONY: install-dev build dev fmt apidoc