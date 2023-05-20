clean:
	go fmt ./cmd/*
	go fmt ./internal/*

run: 
	go run ./cmd/web

test:
	go test -v ./cmd/web

all: clean test run
