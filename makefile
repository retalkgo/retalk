APP_NAME = Retalk
VERSION ?= $(shell git describe --abbrev=0 --tags)
COMMIT_HASH ?= $(shell git log --pretty=format:"%h" -1)
COMMON_LDFLAGS = -ldflags " -X retalk/internal/version.Version=${VERSION} \
							-X retalk/internal/version.CommitHash=${COMMIT_HASH} \
							-s -w"
BIN = bin/retalk

all: install build

install:
	@echo "Retalk ${VERSION}-${COMMIT_HASH} installing..."
	@go mod tidy

fmt:
	@echo "Retalk dev-${COMMIT_HASH} formating..."
	@gofmt -w .

fmt-swagger:
	@echo "Retalk dev-${COMMIT_HASH} formating swagger..."
	@swag fmt

update-swagger: fmt-swagger
	@echo "Retalk ${VERSION}-${COMMIT_HASH} updating swagger..."
	@swag init -g server/server.go

gen:
	@echo "Retalk ${VERSION}-${COMMIT_HASH} generating code..."
	@go run . gen

dev-build: gen
	@echo "Retalk ${VERSION}-${COMMIT_HASH} dev building..."
	@go build -o ${BIN} ${COMMON_LDFLAGS}

dev-run: dev-build
	@echo "Retalk ${VERSION}-${COMMIT_HASH} dev running..."
	@${BIN} start

build: gen update-swagger
	@echo "Retalk ${VERSION}-${COMMIT_HASH} production building..."
	@go build -o ${BIN} ${COMMON_LDFLAGS}