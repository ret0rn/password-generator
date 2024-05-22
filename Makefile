build:
	go mod download
	go build -o ./build/pass-gen ./cmd/main/main.go
run:
	go mod download
	go build -o ./build/pass-gen ./cmd/main/main.go
	./build/pass-gen

.DEFAULT_GOAL := run