IMAGE_REGISTRY 		?= neketsky
IMAGE_NAME 			?= gg-test
FULL_IMAGE_NAME		:= $(IMAGE_REGISTRY)/$(IMAGE_NAME)
VERSION 			?= dev
GO_PACKAGES			:= $(shell go list ./... | grep -v vendor)
GO_FILES        	:= $(shell find . -type f -name '*.go')
REPOSITORY_PATH 	:= gg-test
GO_VERSION 			:= golang:1.20-bullseye
DOCKER_RUNNER 		:= docker run --network host -v $(CURDIR):/go/src/$(REPOSITORY_PATH) -w /go/src/$(REPOSITORY_PATH) ${GO_VERSION}

.PHONY: build
build:
	@docker build -f ./build/Dockerfile --build-arg APP_VERSION=$(VERSION) -t $(FULL_IMAGE_NAME):$(VERSION) .

.PHONY: test
test: postgres-up
	${DOCKER_RUNNER} go test -v --race -cover -tags testmode $(GO_PACKAGES)

.PHONY: postgres-up
postgres-up:
	docker run --rm -d --name postgres -p 5432:5432 -e POSTGRES_HOST_AUTH_METHOD=trust postgres

.PHONY: postgres-down
postgres-down:
	docker rm -f postgres

.PHONY: migrate-up
migrate-up:
	docker run -v $(PWD)/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable up

.PHONY: migrate-down
migrate-down:
	docker run -v $(PWD)/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable down -all

.PHONY: clean
clean: postgres-down
	@docker rmi -f $(shell docker images -q $(FULL_IMAGE_NAME)) || true

.PHONY: fmt
fmt:
	@go fmt $(GO_PACKAGES)

.PHONY: lint
lint:
	@golangci-lint run
	@gocritic check -enableAll ./...
	@gosec ./...

.PHONY: goimports
goimports:
	@docker run --rm -v $(shell pwd):/data cytopia/goimports -w "$(GO_FILES)"

.PHONY: all
all: fmt lint test build
