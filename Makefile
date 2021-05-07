build:
	go build -v ./cmd/server

run:
	go run -v ./cmd/server

.DEFAULT_GOAL := build