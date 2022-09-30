PROJECT_NAME=$(shell head -n 1 go.mod | cut -d'/' -f3)

BASE=$(shell pwd)
GO_BIN=$(BASE)/bin
TEMP_DIR=$(BASE)/temp

name:
	@echo $(PROJECT_NAME)

wire:
	wire ./...

build: wire
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o $(GO_BIN)/server ./cmd/server || exit

test: clean
	mkdir -p ./temp
	go clean -testcache
	go test ./... -count=1 -v -coverprofile=$(TEMP_DIR)/coverage.out
	go tool cover -html=$(TEMP_DIR)/coverage.out -o $(TEMP_DIR)/coverage.html

clean:
	rm -f ./bin/*
	rm -rf ./temp/*

server: wire
	go run ./cmd/server

start: wire
	go run ./cmd/$(PROJECT_NAME)