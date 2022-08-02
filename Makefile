.PHONY: build
build:
	go mod download
	go build -o ./build ./cmd/main/main.go

.PHONY: run
run:
	go mod download
	go build -o ./build ./cmd/main/main.go
	/build

.DEFAULT_GOAL := run