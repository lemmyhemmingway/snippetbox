clean:
	go fmt ./cmd/*
	go fmt ./internal/*

run: 
	go run ./cmd/web

all: clean run
