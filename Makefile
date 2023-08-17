APP_NAME = Retalk
VERSION ?= $(shell git describe --abbrev=0 --tags)
ifeq ($(VERSION),)
VERSION = dev
endif
COMMIT_HASH ?= $(shell git log --pretty=format:"%h" -1)
COMMON_LDFLAGS = -ldflags " -X github.com/retalkgo/retalk/internal/version.Version=${VERSION} \
							-X github.com/retalkgo/retalk/internal/version.CommitHash=${COMMIT_HASH} \
							-s -w"
BIN = bin/retalk

all: install build

install:
	@echo "Retalk ${VERSION}-${COMMIT_HASH} installing..."
	go mod tidy

fmt:
	@echo "Retalk dev-${COMMIT_HASH} formating..."
	gofumpt -l -w -extra .

fmt-swagger:
	@echo "Retalk dev-${COMMIT_HASH} formating swagger..."
	swag fmt

build-apidoc: fmt-swagger update-swagger
	@echo "Retalk dev-${COMMIT_HASH} apidoc building..."
	pnpm dlx @redocly/cli build-docs docs/swagger.yaml -o apidoc/index.html

serve-apidoc: build-apidoc
	@echo "Retalk dev-${COMMIT_HASH} apidoc serving..."
	pnpm dlx serve apidoc

build-frontend:
	@echo "Retalk Frontend ${VERSION}-${COMMIT_HASH} production building..."
	pnpm build

build-docker:
	@echo "Retalk ${VERSION}-${COMMIT_HASH} for docker production building..."
	docker build -t ghcr.io/retalkgo/retalk:${VERSION} .

update-swagger: fmt-swagger
	@echo "Retalk ${VERSION}-${COMMIT_HASH} updating swagger..."
	swag init -g server/server.go

gen:
	@echo "Retalk ${VERSION}-${COMMIT_HASH} generating code..."
	go run . gen

dev-build: update-swagger fmt build-apidoc
	@echo "Retalk dev-${COMMIT_HASH} dev building..."
	go build -o ${BIN} ${COMMON_LDFLAGS}

dev-run: dev-build
	@echo "Retalk dev-${COMMIT_HASH} dev running..."
	${BIN} start

build: gen update-swagger fmt build-apidoc build-frontend
	@echo "Retalk ${VERSION}-${COMMIT_HASH} production building..."
	go build -o ${BIN} ${COMMON_LDFLAGS}

